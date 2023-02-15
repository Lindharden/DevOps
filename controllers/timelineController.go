package controllers

import (
	globals "DevOps/globals"
	helpers "DevOps/helpers"
	model "DevOps/model"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const PAGE_SIZE = 30

func UserTimelineHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey).(model.User)
		db := helpers.GetTypedDb(c)

		userProfileName := c.Param("username")
		following := false
		fmt.Print(userProfileName)
		var profiles =  []model.User{}
		db.Select(&profiles,`select * from user where username = ?`, userProfileName);
	

		if(!helpers.RequestedUserExists(userProfileName, profiles)) {
			 c.JSON(http.StatusNotFound, gin.H{"error": "User does not exist"})
			 return
		} 
	
		entries := []model.TimelineMessage{}
		db.Select(&entries, `select message.*, user.* from message, user
        where message.flagged = 0 and message.author_id = ?
        order by message.pub_date desc limit ?`, profiles[0].UserId, PAGE_SIZE)
		// user timeline
		// gin.H should contain a title text + user object
		c.HTML(http.StatusOK, "timeline.html", gin.H{
			"user":         user,
			"user_profile": profiles[0],
			"followed":     following,
			"isSelf":       profiles[0].UserId == user.UserId,
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
