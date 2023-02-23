package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	helpers "DevOps/helpers"
	"DevOps/model"
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

func RegisterPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := helpers.GetTypedDb(c)
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

		if helpers.EmptyUserPass(username, password) {
			c.HTML(http.StatusBadRequest, "register.html", gin.H{"content": "You have to enter a value"})
			return
		}

		if !helpers.CheckUserPasswords(password, password2) {
			c.HTML(http.StatusBadRequest, "register.html", gin.H{"content": "The two passwords do not match"})
			return
		}

		if !helpers.CheckUserEmail(email) {
			c.HTML(http.StatusBadRequest, "register.html", gin.H{"content": "You have to enter a valid email address"})
			return
		}

		if helpers.CheckUsernameExists(db, username) {
			c.HTML(http.StatusBadRequest, "register.html", gin.H{"content": "The username is already taken"})
			return
		}

		pw_hash, err := helpers.HashPassword(password)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
		}
		db.Exec("insert into user (username, email, pw_hash) values (?, ?, ?)", username, email, pw_hash)

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
		db := helpers.GetTypedDb(c)

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
		userStruct := model.User{}
		db.Get(&userStruct, `select * from user where username = ?`, username)
		log.Println("heya", userStruct)
		if err := helpers.SetUserSession(c, userStruct); err != nil {
			c.HTML(http.StatusInternalServerError, "login.html", gin.H{"content": "Failed to save session"})
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/public")
	}
}

func LogoutGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := helpers.GetUserSession(c); err != nil {
			log.Println("Invalid session token")
			return
		}

		if err := helpers.TerminateUserSession(c); err != nil {
			log.Println("Failed to save session:", err)
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/login")
	}
}
