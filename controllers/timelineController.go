package controllers

import (
	globals "DevOps/globals"
	helpers "DevOps/helpers"
	model "DevOps/model"
	"net/http"
	"time"

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

		//get the requested user
		var profile = model.User{}
		user_exists_err := db.Get(&profile, `select * from user where username = ?`, userProfileName)

		if user_exists_err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User does not exist"})
			return
		}

		//If the user is signed in, check if we follow said user or is that user ourselves
		if user != nil {
			var following interface{}
			err := db.Get(&following, `select 1 from follower where
            follower.who_id = ? and follower.whom_id = ?`, user.(model.User).UserId, profile.UserId)
			//error will be nil if zero rows are returned
			following = err != nil
			isSelf = user.(model.User).UserId == profile.UserId
		}

		//get all the messages from the requested user
		entries := []model.TimelineMessage{}
		db.Select(&entries, `select message.*, user.* from message, user
        where message.flagged = 0 and message.author_id = ?
        order by message.pub_date desc limit ?`, profile.UserId, PAGE_SIZE)

		c.HTML(http.StatusOK, "timeline.html", gin.H{
			"user":         user,
			"user_profile": profile,
			"followed":     following,
			"isSelf":       isSelf,
			"messages":     entries,
			"title":        "Timeline",
		})
	}
}

func PublicTimelineHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := helpers.GetTypedDb(c)
		entries := []model.TimelineMessage{}
		db.Select(&entries, `select message.*, user.* from message, user
        where message.flagged = 0 and message.author_id = user.user_id
        order by message.pub_date desc limit ?`, PAGE_SIZE)
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

func SelfTimeline() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := helpers.GetTypedDb(c)
		session := sessions.Default(c)
		user := session.Get(globals.Userkey).(model.User)
		timelineEntries := []model.TimelineMessage{}
		db.Select(&timelineEntries, `select message.*, user.* from message, user
        where message.flagged = 0 and message.author_id = user.user_id and (
            user.user_id = ? or
            user.user_id in (select whom_id from follower
                                    where who_id = ?))
        order by message.pub_date desc limit ?`, user.UserId, user.UserId, PAGE_SIZE)

		c.HTML(http.StatusOK, "timeline.html", gin.H{
			"user":         user,
			"user_profile": nil,
			"followed":     false,
			"isSelf":       true,
			"messages":     timelineEntries,
			"title":        `Timeline`,
		})
	}
}

func AddMessageHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey).(model.User)
		text := c.PostForm("text")

		if text != "" {
			db := helpers.GetTypedDb(c)
			db.Exec(`insert into message (author_id, text, pub_date, flagged)
            values (?, ?, ?, 0)`, user.UserId, text, time.Now().Unix())
		}
	}
}
