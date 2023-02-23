package middleware

import (
	"github.com/gin-gonic/gin"

	"log"
	"net/http"

	"DevOps/helpers"
)

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
