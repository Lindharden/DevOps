package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"DevOps/globals"
	helpers "DevOps/helpers"
	model "DevOps/model/gorm"
	simModels "DevOps/model/simulatorModel"
)

func FollowHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err0 := helpers.GetUserSession(c)
		if err0 != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		username := c.Param("username") // name of user to follow
		action := c.Param("action")     // follow or unfollow

		db := globals.GetGormDatabase()
		whom_id, err := helpers.GetUserIdGorm(db, username)

		if err != nil {
			c.AbortWithStatus(404)
		} else {
			if action == "/follow" {
				db.Create(&model.Following{UserID: user.ID, WhomId: whom_id})
			} else if action == "/unfollow" {
				db.Where(&model.Following{UserID: user.ID, WhomId: whom_id}).Unscoped().Delete(&model.Following{})
			}
		}

		c.Redirect(http.StatusMovedPermanently, "/"+username)
	}
}

func SimFollowHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request simModels.FollowRequest

		if err := c.BindJSON(&request); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		db := globals.GetGormDatabase()

		username := c.Param("username")
		userId, err := helpers.GetUserIdGorm(db, username)
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

		targetUserId, err := helpers.GetUserIdGorm(db, targetUsername)
		if err != nil {
			// TODO: This has to be another error, likely 500???
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		if request.Follow != "" {
			db.Create(&model.Following{UserID: userId, WhomId: targetUserId})
		} else if request.Unfollow != "" {
			db.Where(&model.Following{UserID: userId, WhomId: targetUserId}).Unscoped().Delete(&model.Following{})
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

		db := globals.GetGormDatabase()

		username := c.Param("username")
		_, err = helpers.GetUserIdGorm(db, username)

		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		var list model.User

		db.Preload("Followings", func(tx *gorm.DB) *gorm.DB {
			return tx.Limit(numberFollowers)
		}).
			Where(&model.User{Username: username}).
			Preload("Followings.WhomUser").
			First(&list)

		c.JSON(http.StatusOK, gin.H{
			"follows": helpers.Map(list.Followings, func(x model.Following) string { return x.WhomUser.Username }),
		})
	}
}
