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

func UserTimelineHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		db := helpers.GetTypedDb(c)

		userProfileName := c.Param("username")
		following := false
		isSelf := false

		
		var profile =  model.User{}
		user_exists_err := db.Get(&profile,`select * from user where username = ?`, userProfileName);
		

		if(user_exists_err != nil) {
			 c.JSON(http.StatusNotFound, gin.H{"error": "User does not exist"})
			 return
		} 

		if(user != nil) {
			var following interface {}
			err := db.Get(&following, `select 1 from follower where
            follower.who_id = ? and follower.whom_id = ?`,user.(model.User).UserId, profile.UserId)
			following = err != nil
			isSelf = user.(model.User).UserId == profile.UserId
		}
	
		entries := []model.TimelineMessage{}
		db.Select(&entries, `select message.*, user.* from message, user
        where message.flagged = 0 and message.author_id = ?
        order by message.pub_date desc limit ?`, profile.UserId, PAGE_SIZE)
		// user timeline
		// gin.H should contain a title text + user object
		c.HTML(http.StatusOK, "timeline.html", gin.H{
			"user":         user,
			"user_profile": profile,
			"followed":     following,
			"isSelf":       isSelf,
			"messages":     entries,
			"title":        "Timeline",
		})

		// should contain a user_profile object for a user timeline
		// should contain whether following or not
		// should contain whether it is yourself
	}
}


func PublicTimelineHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
	db := helpers.GetTypedDb(c)
	entries := []model.TimelineMessage{}
		db.Select(&entries, `select message.*, user.* from message, user
        where message.flagged = 0 and message.author_id = user.user_id
        order by message.pub_date desc limit ?`, PAGE_SIZE)
		// user timeline
		// gin.H should contain a title text + user object
		c.HTML(http.StatusOK, "timeline.html", gin.H{
			"user":         nil,
			"user_profile": nil,
			"followed":     false,
			"isSelf":       false,
			"messages":     entries,
			"title":        "Timeline",
		})
	}
}

func PrivateTimelineHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// if not signed in, reroute to public timeline
	}
}

func AddMessageHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.HTML(http.StatusBadRequest, "login.html", gin.H{"content": "You are not logged in"})
			return
		}
	}
}
