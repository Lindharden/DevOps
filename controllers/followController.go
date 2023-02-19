package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	globals "DevOps/globals"
	helpers "DevOps/helpers"
	"DevOps/model"
)

func FollowHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey).(model.User)
		username := c.Param("username") // name of user to follow
		action := c.Param("action")     // follow or unfollow

		db := helpers.GetTypedDb(c)
		whom_id, err := helpers.GetUserId(db, username)

		if err != nil {
			c.AbortWithStatus(404)
		} else {
			if action == "/follow" {
				db.Exec("insert into follower (who_id, whom_id) values (?, ?)", user.UserId, whom_id)
			} else if action == "/unfollow" {
				db.Exec("delete from follower where who_id=? and whom_id=?", user.UserId, whom_id)
			}
			c.AbortWithStatus(200)
		}
	}
}
