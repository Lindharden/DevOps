package controllers

import (
	"github.com/gin-contrib/sessions"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	globals "DevOps/globals"
	helpers "DevOps/helpers"
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
		session := sessions.Default(c)
		session.Set(globals.LatestRequest, c.Request) // save latest Register Post request (used in Sim)
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
			c.HTML(http.StatusBadRequest, "register.html", gin.H{"content": "Parameters can't be empty"})
			return
		}

		if !helpers.CheckUserRegisterInfo(username, password, password2, email) {
			c.HTML(http.StatusBadRequest, "register.html", gin.H{"content": "Parameters are bad"})
			// TODO: make specific messages for bad email, unequal pw1/pw2
			return
		}

		session.Set(globals.Userkey, username)
		if err := session.Save(); err != nil {
			c.HTML(http.StatusInternalServerError, "register.html", gin.H{"content": "Failed to save session"})
			return
		}

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
			"user":    user,
		})
	}
}

func LoginPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.HTML(http.StatusBadRequest, "login.html", gin.H{"content": "Please logout first"})
			return
		}

		username := c.PostForm("username")
		password := c.PostForm("password")

		if helpers.EmptyUserPass(username, password) {
			c.HTML(http.StatusBadRequest, "login.html", gin.H{"content": "Parameters can't be empty"})
			return
		}

		if !helpers.CheckUserPass(username, password) {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{"content": "Incorrect username or password"})
			return
		}

		session.Set(globals.Userkey, username)
		if err := session.Save(); err != nil {
			c.HTML(http.StatusInternalServerError, "login.html", gin.H{"content": "Failed to save session"})
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/index")
	}
}

func LogoutGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		log.Println("logging out user:", user)
		if user == nil {
			log.Println("Invalid session token")
			return
		}
		session.Delete(globals.Userkey)
		if err := session.Save(); err != nil {
			log.Println("Failed to save session:", err)
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/")
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
