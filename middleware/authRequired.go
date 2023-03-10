package middleware

import (
	"github.com/gin-gonic/gin"

	"log"
	"net/http"

	"DevOps/helpers"
)

func CheckRequestFromSimulator(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header != "Basic c2ltdWxhdG9yOnN1cGVyX3NhZmUh" {
		c.JSON(http.StatusNotFound, gin.H{"error": "You are not authorized to use this resource!"})
		c.Abort()
		return
	}
	c.Next()
}

func AuthRequired(c *gin.Context) {
	_, err := helpers.GetUserSession(c)
	if err != nil {
		log.Println("User not logged in")
		c.Redirect(http.StatusMovedPermanently, "/login")
		c.Abort()
		return
	}
	c.Next()
}
