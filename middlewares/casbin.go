package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gostars/service"
	"gostars/utils"
	"gostars/utils/code"
	"net/http"
)

var casbinService = new(service.CasbinService)

func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		//claims, _ := c.Get("claims")
		//waitUse := claims.(*request.CustomClaims)

		obj := c.Request.URL.RequestURI()
		act := c.Request.Method
		sub := c.Query("authorityid")
		fmt.Println(sub, obj, act)
		e := casbinService.Casbin()
		msg, _ := e.Enforce(sub, obj, act)
		fmt.Println(msg)

		// debug env
		if msg || utils.AppMode == "debug" {
			c.Next()
			//if msg {
			//	c.Next()
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": code.ErrorUserNoRight,
				"msg":    code.GetErrMsg(code.ErrorUserNoRight),
			})
			c.Abort()
			return
		}
	}
}
