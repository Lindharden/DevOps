package controllers

import (
	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
	"net/http"

	globals "DevOps/globals"
)

func FollowHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		//username := c.Param("username") // name of user to follow
		//action := c.Param("action") // follow or unfollow
		if user != nil {
			c.HTML(http.StatusBadRequest, "login.html", gin.H{"content": "You are not logged in"})
			return
		}
		// TODO: Check that username exists, and execute follow/unfollow
	}
}