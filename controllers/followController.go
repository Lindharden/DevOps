package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	helpers "DevOps/helpers"
)

func FollowHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err0 := helpers.GetUserSession(c)
		if err0 != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
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
