package controllers

import (
	globals "DevOps/globals"
	helpers "DevOps/helpers"
	model "DevOps/model/gorm"
	"net/http"
	"strconv"
	"time"

	simModels "DevOps/model/simulatorModel"

	"github.com/gin-gonic/gin"
)

func AddMessageHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		helpers.TweetsProcessed.Inc()
		user, err := helpers.GetUserSession(c)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		text := c.PostForm("text")

		if text != "" {
			db := globals.GetGormDatabase()
			db.Create(&model.Message{UserID: user.ID,
				User:    model.User{},
				Text:    text,
				PubDate: time.Now().Unix(),
				Flagged: 0})
		}

		c.Redirect(http.StatusMovedPermanently, "/public")
	}
}

func GetMessageHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// query db
		db := globals.GetGormDatabase()
		entries := []model.Message{}

		// check for parameter "no" (number of messages)
		noMsgs, err := strconv.Atoi(c.Query("no"))
		if err != nil {
			// if undefined, use default value
			noMsgs = 100
		}

		db.Preload("User").
			Where(&model.Message{Flagged: 0}).
			Order("pub_date desc").
			Limit(noMsgs).
			Find(&entries)

		// filter messages
		var messageList []simModels.FilteredMessageRequest
		for _, message := range entries {
			messageList = append(messageList,
				simModels.FilteredMessageRequest{
					Text:     message.Text,
					PubDate:  message.PubDate,
					Username: message.User.Username,
				})
		}

		c.JSON(200, messageList)
	}
}

func GetMessageUserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := globals.GetGormDatabase()

		// convert username to user id
		username := c.Param(globals.Username)
		user_id, err := helpers.GetUserIdGorm(db, username)

		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		// check for parameter "no" (number of messages)
		noMsgs, err := strconv.Atoi(c.Query("no"))
		if err != nil {
			// if undefined, use default value
			noMsgs = 100
		}

		entries := []model.Message{}

		db.Preload("User").
			Where(&model.Message{Flagged: 0, UserID: user_id}).
			Order("pub_date desc").
			Limit(noMsgs).
			Find(&entries)

		// filter messages
		var messageList []simModels.FilteredMessageRequest
		for _, message := range entries {
			messageList = append(messageList,
				simModels.FilteredMessageRequest{
					Text:     message.Text,
					PubDate:  message.PubDate,
					Username: message.User.Username,
				})
		}

		c.JSON(200, messageList)
	}
}

func PostMessageUserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		helpers.TweetsProcessed.Inc()
		var postMessage simModels.MessageRequest

		// bind JSON body to postMessage
		if err := c.BindJSON(&postMessage); err != nil {
			globals.GetLogger().Errorw("Bad post message request", "error", err.Error())
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// get DB
		db := globals.GetGormDatabase()

		// get username, and convert to user id
		username := c.Param("username")
		userId, err := helpers.GetUserIdGorm(db, username)
		if err != nil {
			c.AbortWithStatus(404)
			return
		}

		time := time.Now().Unix()

		db.Create(&model.Message{
			UserID:  userId,
			Text:    postMessage.Content,
			PubDate: time,
			Flagged: 0})

		// exit with status 204
		c.Status(http.StatusNoContent)
	}
}
