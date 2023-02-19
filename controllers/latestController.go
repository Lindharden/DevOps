package controllers

import (
	"DevOps/globals"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LatestGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// return latest request
		session := sessions.Default(c)
		reqJson := session.Get(globals.Latest).([]byte)
		c.JSON(200, reqJson)
	}
}
