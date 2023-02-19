package controllers

import (
	globals "DevOps/globals"
	helpers "DevOps/helpers"
	model "DevOps/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const PAGE_SIZE = 30

func PublicTimelineHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		db := helpers.GetTypedDb(c)
		// userProfile := c.Param("username")
		entries := []model.TimelineMessage{}
		db.Select(&entries, `select message.*, user.* from message, user
        where message.flagged = 0 and message.author_id = user.user_id
        order by message.pub_date desc limit ?`, PAGE_SIZE)
		// user timeline
		// gin.H should contain a title text + user object
		c.HTML(http.StatusOK, "timeline.html", gin.H{
			"user":         user,
			"user_profile": nil,
			"followed":     false,
			"isSelf":       false,
			"messages":     entries,
			"title":        "Timeline",
		})

		// should contain a user_profile object for a user timeline
		// should contain whether following or not
		// should contain whether it is yourself
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

func GetMessageHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		globals.SaveRequest(c)

		db := helpers.GetTypedDb(c)
		entries := []model.TimelineMessage{}
		db.Select(&entries, `SELECT message.*, user.* FROM message, user
        WHERE message.flagged = 0 AND message.author_id = user.user_id
        ORDER BY message.pub_date DESC LIMIT ?`, PAGE_SIZE)

		type MessageList struct {
			Messages []model.TimelineMessage `json:"messages"`
		}

		messageList := MessageList{Messages: entries}

		jsonBytes, err := json.Marshal(messageList)
		if err != nil {
			log.Println("Error encoding JSON:", err)
			return
		}

		c.JSON(200, jsonBytes)
	}
}
