package controllers

import (
	"DevOps/globals"
	helpers "DevOps/helpers"
	gormModel "DevOps/model/gorm"
	"net/http"

	"github.com/gin-gonic/gin"
)

const PAGE_SIZE = 30

func UserTimelineHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := helpers.GetUserSession(c)
		db := globals.GetGormDatabase()

		userProfileName := c.Param("username")
		following := false
		isSelf := false
		isAuthorized := false

		//get the requested user
		var profile = gormModel.User{}
		user_exists_err := db.Where(gormModel.User{Username: userProfileName}).First(&profile)

		if user_exists_err.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User does not exist"})
			return
		}

		//If the user is signed in, check if we follow said user or is that user ourselves
		if err == nil {
			var result = gormModel.Following{}
			res := db.Where(gormModel.Following{WhoId: user.ID, WhomId: profile.ID}).Limit(1).First(&result)

			//error will be nil if zero rows are returned
			following = res.Error == nil
			isSelf = user.ID == profile.ID
			isAuthorized = true
		}

		//get all the messages from the requested user
		entries := []gormModel.Message{}

		db.Preload("User").
			Where(gormModel.Message{Flagged: 0, UserID: profile.ID}).
			Order("pub_date desc").
			Limit(PAGE_SIZE).Find(&entries)

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
		db := globals.GetGormDatabase()
		user, err := helpers.GetUserSession(c)
		isAuthorized := false
		if err == nil {
			isAuthorized = true
		}
		var entries []gormModel.Message
		db.Preload("User").
			Where(gormModel.Message{Flagged: 0}).
			Order("pub_date desc").
			Limit(PAGE_SIZE).
			Find(&entries)

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
		db := globals.GetGormDatabase()
		user, err := helpers.GetUserSession(c)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		timelineEntries := []gormModel.Message{}
		subQuery := db.Select("whom_id").Where(&gormModel.Following{WhoId: user.ID}).Table("followings")

		db.Preload("User").
			Where(&gormModel.Message{UserID: user.ID}).
			Or("user_id in (?)", subQuery).
			Order("pub_date desc").
			Limit(PAGE_SIZE).
			Find(&timelineEntries)

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
