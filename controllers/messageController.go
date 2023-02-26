package controllers

import (
	globals "DevOps/globals"
	helpers "DevOps/helpers"
	model "DevOps/model"
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
			db := globals.DB
			db.Exec(`insert into message (author_id, text, pub_date, flagged)
            values (?, ?, ?, 0)`, user.UserId, text, time.Now().Unix())
		}

		c.Redirect(http.StatusMovedPermanently, "/public")
	}
}

func GetMessageHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// query db
		db := globals.DB
		entries := []model.TimelineMessage{}

		// check for parameter "no" (number of messages)
		noMsgs, err := strconv.Atoi(c.Query("no"))
		if err != nil {
			// if undefined, use default value
			noMsgs = 100
		}
		db.Select(&entries, `SELECT message.*, user.* FROM message, user
        WHERE message.flagged = 0 AND message.author_id = user.user_id
        ORDER BY message.pub_date DESC LIMIT ?`, noMsgs)

		// filter messages
		var messageList []simModels.FilteredMessageRequest
		for _, message := range entries {
			messageList = append(messageList,
				simModels.FilteredMessageRequest{
					Text:     message.Text,
					PubDate:  message.PubDate,
					Username: message.Username,
				})
		}

		c.JSON(200, messageList)
	}
}

func GetMessageUserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := globals.DB

		// convert username to user id
		username := c.Param(globals.Username)
		user_id, err := helpers.GetUserId(db, username)

		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		entries := []model.TimelineMessage{}
		db.Select(&entries, `SELECT message.*, user.* FROM message, user 
		WHERE message.flagged = 0 AND
		user.user_id = message.author_id AND user.user_id = ?
		ORDER BY message.pub_date DESC LIMIT 100`, user_id)

		// filter messages
		var messageList []simModels.FilteredMessageRequest
		for _, message := range entries {
			messageList = append(messageList,
				simModels.FilteredMessageRequest{
					Text:     message.Text,
					PubDate:  message.PubDate,
					Username: message.Username,
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
		db := globals.DB

		// get username, and convert to user id
		username := c.Param("username")
		userId, err := helpers.GetUserId(db, username)
		if err != nil {
			c.AbortWithStatus(404)
		}

		time := time.Now().Unix()

		db.Exec(`insert into message (author_id, text, pub_date, flagged) values (?, ?, ?, 0)`,
			userId, postMessage.Content, time)

		// exit with status 204
		c.Status(http.StatusNoContent)
	}
}
