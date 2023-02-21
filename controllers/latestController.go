package controllers

import (
	"DevOps/globals"

	"github.com/gin-gonic/gin"
)

func LatestGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		c.JSON(200, gin.H{"latest": globals.GetLatestRequestId()})
	}
}
