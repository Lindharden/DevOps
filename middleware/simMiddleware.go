package middleware

import (
	"DevOps/globals"
	"strconv"

	"github.com/gin-gonic/gin"
)



func SimulationRequest(c *gin.Context){
	latestRequestId, err :=  strconv.Atoi(c.Query("latest"))
	if err == nil && latestRequestId != 0 {
		globals.SetLatestRequestId(latestRequestId)
	}
}