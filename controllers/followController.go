package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	globals "DevOps/globals"
	helpers "DevOps/helpers"
	model "DevOps/model"
	simModels "DevOps/model/simulatorModel"
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

func SimFollowHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request simModels.FollowRequest

		if err := c.BindJSON(&request); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		db := helpers.GetTypedDb(c)

		username := c.Param("username")
		userId, err := helpers.GetUserId(db, username)
		if err != nil {
			// TODO: This has to be another error, likely 500???
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		var targetUsername string
		if request.Follow != "" {
			targetUsername = request.Follow
		} else {
			targetUsername = request.Unfollow
		}

		targetUserId, err := helpers.GetUserId(db, targetUsername)
		if err != nil {
			// TODO: This has to be another error, likely 500???
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		if request.Follow != "" {
			db.Exec("insert into follower (who_id, whom_id) values (?, ?)", userId, targetUserId)
		} else if request.Unfollow != "" {
			db.Exec("DELETE FROM follower WHERE who_id=? and WHOM_ID=?", userId, targetUserId)
		}

		c.Status(http.StatusNoContent)
	}
}

func SimGetFollowHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		numberFollowers, err := strconv.Atoi(c.Query("no"))

		if err != nil {
			numberFollowers = 100
		}

		db := helpers.GetTypedDb(c)

		query := `SELECT user.username FROM user
		INNER JOIN follower ON follower.whom_id=user.user_id
		WHERE follower.who_id=?
		LIMIT ?`

		username := c.Param("username")
		userId, err := helpers.GetUserId(db, username)

		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		var list []string
		db.Select(&list, query, userId, numberFollowers)

		c.JSON(http.StatusOK, gin.H{
			"follows": list,
		})
	}
}
