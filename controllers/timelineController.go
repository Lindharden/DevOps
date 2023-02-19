package controllers

import (
	globals "DevOps/globals"
	helpers "DevOps/helpers"
	model "DevOps/model"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const PAGE_SIZE = 30

func PublicTimelineHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		db := helpers.GetTypedDb(c)
		// userProfile := c.Param("username")
		entries := []model.TimelineMessage{}
		db.Select(&entries, `select message.*, user.* from message, user
        where message.flagged = 0 and message.author_id = user.user_id
        order by message.pub_date desc limit ?`, PAGE_SIZE)
		// user timeline
		// gin.H should contain a title text + user object
		c.HTML(http.StatusOK, "timeline.html", gin.H{
			"user":         user,
			"user_profile": nil,
			"followed":     false,
			"isSelf":       false,
			"messages":     entries,
			"title":        "Timeline",
		})

		// should contain a user_profile object for a user timeline
		// should contain whether following or not
		// should contain whether it is yourself
	}
}

func PrivateTimelineHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// if not signed in, reroute to public timeline
	}
}
