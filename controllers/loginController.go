package controllers

import (
	"github.com/gin-contrib/sessions"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	globals "DevOps/globals"
	helpers "DevOps/helpers"
	"DevOps/model"
)

func RegisterGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.HTML(http.StatusBadRequest, "register.html",
				gin.H{
					"content": "Please logout first",
					"user":    user,
				})
			return
		}
		c.HTML(http.StatusOK, "register.html", gin.H{
			"content": "",
			"user":    user,
		})
	}
}

func RegisterPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := helpers.GetTypedDb(c)
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)

		if user != nil {
			c.HTML(http.StatusBadRequest, "register.html", gin.H{"content": "Please logout first"})
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
			c.AbortWithStatus(404)
		}
		db.Exec("insert into user (username, email, pw_hash) values (?, ?, ?)", username, email, pw_hash)

		c.Redirect(http.StatusMovedPermanently, "/login")
	}
}

func LoginGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.HTML(http.StatusBadRequest, "login.html",
				gin.H{
					"content": "Please logout first",
					"user":    user,
				})
			return
		}
		c.HTML(http.StatusOK, "login.html", gin.H{
			"content": "",
		})
	}
}

func LoginPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		db := helpers.GetTypedDb(c)
		_, err := helpers.GetUserSession(c)
		if err == nil {
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
		db.Select(&userStruct, "select * from user where username = ?", username)

		helpers.SetUserSession(c, userStruct)
		if err := session.Save(); err != nil {
			c.HTML(http.StatusInternalServerError, "login.html", gin.H{"content": "Failed to save session"})
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/public")
	}
}

func LogoutGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user == nil {
			log.Println("Invalid session token")
			return
		}
		session.Delete(globals.Userkey)
		if err := session.Save(); err != nil {
			log.Println("Failed to save session:", err)
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/login")
	}
}

func IndexGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content": "This is an index page...",
			"user":    user,
		})
	}
}
