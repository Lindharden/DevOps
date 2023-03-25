package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"DevOps/globals"
	helpers "DevOps/helpers"
	gormModel "DevOps/model/gorm"
	simModels "DevOps/model/simulatorModel"
)

func RegisterGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := helpers.GetUserSession(c)
		if err == nil {
			c.HTML(http.StatusBadRequest, "register.html",
				gin.H{
					"content": "Please logout first",
					"user":    user,
				})
			return
		}
		c.HTML(http.StatusOK, "register.html", gin.H{
			"content": "",
			"user":    nil,
		})
	}
}

func SimRegisterPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := globals.GetGormDatabase()

		var registerData simModels.RegisterRequest

		if err := c.BindJSON(&registerData); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		_, err := helpers.RegisterUser(db, registerData.Username, registerData.Pwd, registerData.Pwd, registerData.Email)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error_msg": err.Error()})
			return
		}

		c.AbortWithStatus(http.StatusNoContent)

	}
}

func RegisterPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := globals.GetGormDatabase()
		user, err := helpers.GetUserSession(c)
		if err == nil {
			c.HTML(http.StatusBadRequest, "register.html",
				gin.H{
					"content": "Please logout first",
					"user":    user,
				})
			return
		}

		username := c.PostForm("username")
		password := c.PostForm("password")
		password2 := c.PostForm("password2")
		email := c.PostForm("email")

		_, signUpError := helpers.RegisterUser(db, username, password, password2, email)

		if signUpError != nil {
			c.HTML(http.StatusBadRequest, "register.html", gin.H{"content": signUpError.Error()})
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/login")
	}
}

func LoginGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := helpers.GetUserSession(c)
		if err == nil {
			c.HTML(http.StatusBadRequest, "login.html",
				gin.H{
					"content": "Please logout first",
					"user":    user,
				})
			return
		}
		c.HTML(http.StatusOK, "login.html", nil)
	}
}

func LoginPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := globals.GetGormDatabase()

		if _, err := helpers.GetUserSession(c); err == nil {
			c.HTML(http.StatusBadRequest, "login.html", gin.H{"content": "Please logout first"})
			return
		}

		username := c.PostForm("username")
		password := c.PostForm("password")

		if helpers.EmptyUserPass(username, password) {
			c.HTML(http.StatusBadRequest, "login.html", gin.H{"content": "Parameters can't be empty"})
			return
		}

		if !helpers.CheckUsernameExists(db, username) {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{"content": "Invalid username"})
			return
		}

		if !helpers.ValidatePassword(db, username, password) {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{"content": "Invalid password"})
			return
		}

		var userStruct gormModel.User
		db.Where(&gormModel.User{Username: username}).First(&userStruct)

		if err := helpers.SetUserSession(c, userStruct); err != nil {
			globals.GetLogger().Errorw("Could not session", "user", userStruct.Username)
			c.HTML(http.StatusInternalServerError, "login.html", gin.H{"content": "Failed to save session"})
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/public")
	}
}

func LogoutGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := helpers.GetUserSession(c); err != nil {
			globals.GetLogger().Warn("Invalid session token")
			return
		}

		if err := helpers.TerminateUserSession(c); err != nil {
			globals.GetLogger().Errorw("Failed to save session", "error", err.Error())
			return
		}

		c.Redirect(http.StatusFound, "/login")
	}
}
