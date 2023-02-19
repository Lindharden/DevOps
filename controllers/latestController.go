package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LatestGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// return latest request
		session := sessions.Default(c)
		reqJson := session.Get("latest").([]byte)
		c.JSON(200, reqJson)
	}
}
