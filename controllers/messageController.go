package controllers

import (
	globals "DevOps/globals"
	helpers "DevOps/helpers"
	model "DevOps/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

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

type filteredMessage struct {
	Text     string `db:"text"`
	PubDate  int64  `db:"pub_date"`
	Username string `db:"username"`
}

type MessageList struct {
	Messages []filteredMessage `json:"messages"`
}

func GetMessageHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		globals.SaveRequest(c)

		// query db
		db := helpers.GetTypedDb(c)
		entries := []model.TimelineMessage{}
		db.Select(&entries, `SELECT message.*, user.* FROM message, user
        WHERE message.flagged = 0 AND message.author_id = user.user_id
        ORDER BY message.pub_date DESC LIMIT 100`)

		// filter messages
		var messageList MessageList = MessageList{}
		for _, message := range entries {
			messageList.Messages = append(messageList.Messages,
				filteredMessage{
					Text:     message.Text,
					PubDate:  message.PubDate,
					Username: message.Username,
				})
		}

		// convert to JSON
		jsonBytes, err := json.Marshal(messageList)
		if err != nil {
			log.Println("Error encoding JSON:", err)
			return
		}

		c.JSON(200, jsonBytes)
	}
}

func GetMessageUserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		globals.SaveRequest(c)

		// convert username to user id
		username := c.Param(globals.Username)
		user_id := globals.GetUserId(username, c)

		db := helpers.GetTypedDb(c)
		entries := []model.TimelineMessage{}
		query := fmt.Sprintf(`SELECT message.*, user.* FROM message, user 
		WHERE message.flagged = 0 AND
		user.user_id = message.author_id AND user.user_id = %d
		ORDER BY message.pub_date DESC LIMIT 100,`, user_id)
		db.Select(&entries, query)

		// filter messages
		var messageList MessageList = MessageList{}
		for _, message := range entries {
			messageList.Messages = append(messageList.Messages,
				filteredMessage{
					Text:     message.Text,
					PubDate:  message.PubDate,
					Username: message.Username,
				})
		}

		// convert to JSON
		jsonBytes, err := json.Marshal(messageList)
		if err != nil {
			log.Println("Error encoding JSON:", err)
			return
		}

		c.JSON(200, jsonBytes)
	}
}
