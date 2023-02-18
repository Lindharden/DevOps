package controllers

import (
	globals "DevOps/globals"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LatestGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// return latest request
		session := sessions.Default(c)
		latest := struct{ latest *http.Request }{latest: session.Get(globals.LatestRequest).(*http.Request)}
		c.JSON(200, latest)
	}
}
