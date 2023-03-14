package middleware

import (
	"DevOps/helpers"
	"time"

	"github.com/gin-gonic/gin"
)

func MeasueRequestTimeMiddleware(c *gin.Context) {
	initialTime := time.Now()
	c.Next()
	requestTime := time.Since(initialTime).Milliseconds()
	helpers.RequestResponseTime.Observe(float64(requestTime))
}
