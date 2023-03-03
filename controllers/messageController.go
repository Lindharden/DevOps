package controllers

import (
	globals "DevOps/globals"
	helpers "DevOps/helpers"
	gormModel "DevOps/model/gorm"
	"net/http"
	"strconv"
	"time"

	simModels "DevOps/model/simulatorModel"

	"github.com/gin-gonic/gin"
)

func AddMessageHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := helpers.GetUserSession(c)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		text := c.PostForm("text")

		if text != "" {
			db := globals.GetGormDatabase()
			db.Create(&gormModel.Message{UserID: user.ID,
				User:    gormModel.User{},
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
		entries := []gormModel.Message{}

		// check for parameter "no" (number of messages)
		noMsgs, err := strconv.Atoi(c.Query("no"))
		if err != nil {
			// if undefined, use default value
			noMsgs = 100
		}

		db.Limit(noMsgs).Select("message.*", "user.*").
			Where("message.flagged = ? AND message.author_id = ?", 0, "user.user_id").
			Order("message.pub_date desc").
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

		entries := []gormModel.Message{}
		db.Limit(noMsgs).Select("message.*", "user.*").
			Where("message.flagged = ? AND user.user_id = ? AND user.user_id = ?", 0, "message.author_id", user_id).
			Order("message.pub_date desc").
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

		var postMessage simModels.MessageRequest

		// bind JSON body to postMessage
		if err := c.BindJSON(&postMessage); err != nil {
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
		}

		time := time.Now().Unix()

		db.Create(&gormModel.Message{
			UserID:  userId,
			Text:    postMessage.Content,
			PubDate: time,
			Flagged: 0})

		// exit with status 204
		c.Status(http.StatusNoContent)
	}
}
