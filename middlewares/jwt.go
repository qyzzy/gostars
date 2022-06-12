package middlewares

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gostars/utils"
	code2 "gostars/utils/code"
	"net/http"
	"strings"
)

type JWT struct {
	JwtKey []byte
}

func NewJwt() *JWT {
	return &JWT{
		[]byte(utils.JwtKey),
	}
}

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var (
	TokenExpired     error = errors.New("token expired")
	TokenNotValidYet error = errors.New("token not valid yet")
	TokenMalformed   error = errors.New("token wrong")
	TokenInvalid     error = errors.New("invalid token")
)

func (j *JWT) CreateToken(claims MyClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.JwtKey)
}

func (j *JWT) ParserToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.JwtKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	}

	return nil, TokenInvalid
}

func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var errCode int
		tokenHeader := c.Request.Header.Get("Authorization")

		if tokenHeader == "" {
			errCode = code2.ErrorTokenExist
			c.JSON(http.StatusOK, gin.H{
				"status":  errCode,
				"message": code2.GetErrMsg(errCode),
			})
			c.Abort()
			return
		}

		checkToken := strings.Split(tokenHeader, " ")
		if len(checkToken) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"status":  errCode,
				"message": code2.GetErrMsg(errCode),
			})
			c.Abort()
			return
		}

		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			c.JSON(http.StatusOK, gin.H{
				"status":  errCode,
				"message": code2.GetErrMsg(errCode),
			})
			c.Abort()
			return
		}

		j := NewJwt()
		// Parse token
		claims, err := j.ParserToken(checkToken[1])
		if err != nil {
			if err == TokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"status":  code2.ERROR,
					"message": "token expired",
					"data":    nil,
				})
				c.Abort()
				return
			}
			// other
			c.JSON(http.StatusOK, gin.H{
				"status":  code2.ERROR,
				"message": err.Error(),
				"data":    nil,
			})
			c.Abort()
			return
		}

		c.Set("username", claims)
		c.Next()
	}
}
