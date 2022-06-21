package middlewares

import (
	"github.com/gin-gonic/gin"
	"gostars/utils/code"
	"gostars/utils/jwt"
	"net/http"
	"strings"
)

func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var errCode int
		tokenHeader := c.Request.Header.Get("Authorization")

		if tokenHeader == "" {
			errCode = code.ErrorTokenExist
			c.JSON(http.StatusOK, gin.H{
				"status":  errCode,
				"message": code.GetErrMsg(errCode),
			})
			c.Abort()
			return
		}

		checkToken := strings.Split(tokenHeader, "Bearer ")
		if len(checkToken) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"status":  errCode,
				"message": code.GetErrMsg(errCode),
			})
			c.Abort()
			return
		}

		if len(checkToken) != 2 {
			c.JSON(http.StatusOK, gin.H{
				"status":  errCode,
				"message": code.GetErrMsg(errCode),
			})
			c.Abort()
			return
		}

		j := jwt.NewJwt()
		// Parse token
		claims, err := j.ParserToken(checkToken[1])
		if err != nil {
			if err == jwt.TokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"status":  code.ERROR,
					"message": "token expired",
					"data":    nil,
				})
				c.Abort()
				return
			}
			// other
			c.JSON(http.StatusOK, gin.H{
				"status":  code.ERROR,
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
