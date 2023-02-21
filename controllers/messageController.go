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

		c.JSON(200, messageList)
	}
}

func GetMessageUserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		globals.SaveRequest(c)
		db := helpers.GetTypedDb(c)

		// convert username to user id
		username := c.Param(globals.Username)
		user_id, err := helpers.GetUserId(db,username)

		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		entries := []model.TimelineMessage{}
		db.Select(&entries, `SELECT message.*, user.* FROM message, user 
		WHERE message.flagged = 0 AND
		user.user_id = message.author_id AND user.user_id = ?
		ORDER BY message.pub_date DESC LIMIT 100`,user_id)

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

		c.JSON(200, messageList)
	}
}
