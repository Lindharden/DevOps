package controllers

import (
	"DevOps/globals"
	helpers "DevOps/helpers"
	model "DevOps/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

const PAGE_SIZE = 30

func UserTimelineHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := helpers.GetUserSession(c)
		db := globals.GetDatabase()

		userProfileName := c.Param("username")
		following := false
		isSelf := false
		isAuthorized := false

		//get the requested user
		var profile = model.User{}
		user_exists_err := db.Get(&profile, `select * from user where username = ?`, userProfileName)

		if user_exists_err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User does not exist"})
			return
		}

		//If the user is signed in, check if we follow said user or is that user ourselves
		if err == nil {
			var result = model.FollowingEntry{}
			err := db.Get(&result, `select * from follower where
            follower.who_id = ? and follower.whom_id = ? limit 1`, user.UserId, profile.UserId)
			//error will be nil if zero rows are returned
			following = err == nil
			isSelf = user.UserId == profile.UserId
			isAuthorized = true
		}

		//get all the messages from the requested user
		entries := []model.TimelineMessage{}
		db.Select(&entries, `select message.*, user.* from message, user
        where message.flagged = 0 and message.author_id = ?
        order by message.pub_date desc limit ?`, profile.UserId, PAGE_SIZE)

		c.HTML(http.StatusOK, "timeline.html", gin.H{
			"authorized":  isAuthorized,
			"user":        user,
			"userProfile": profile,
			"followed":    following,
			"isSelf":      isSelf,
			"messages":    entries,
			"title":       "Timeline",
		})
	}
}

func PublicTimelineHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := globals.GetDatabase()
		user, err := helpers.GetUserSession(c)
		isAuthorized := false
		if err == nil {
			isAuthorized = true
		}
		entries := []model.TimelineMessage{}
		db.Select(&entries, `select message.*, user.* from message, user
        where message.flagged = 0 and message.author_id = user.user_id
        order by message.pub_date desc limit ?`, PAGE_SIZE)
		c.HTML(http.StatusOK, "timeline.html", gin.H{
			"authorized":   isAuthorized,
			"user":         user,
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
		db := globals.GetDatabase()
		user, err := helpers.GetUserSession(c)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		timelineEntries := []model.TimelineMessage{}
		db.Select(&timelineEntries, `select message.*, user.* from message, user
        where message.flagged = 0 and message.author_id = user.user_id and (
            user.user_id = ? or
            user.user_id in (select whom_id from follower
                                    where who_id = ?))
        order by message.pub_date desc limit ?`, user.UserId, user.UserId, PAGE_SIZE)

		c.HTML(http.StatusOK, "timeline.html", gin.H{
			"authorized":   true,
			"user":         user,
			"user_profile": nil,
			"followed":     false,
			"isSelf":       true,
			"messages":     timelineEntries,
			"title":        `Timeline`,
		})
	}
}
