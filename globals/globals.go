package globals

import (
	"encoding/json"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var Secret = []byte("secret")

const Userkey = "user"

const Latest = "latest"

func SaveRequest(c *gin.Context) {
	session := sessions.Default(c)
	reqJson, _ := json.Marshal(c.Request)
	session.Set(Latest, reqJson)
	if err := session.Save(); err != nil {
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{"content": "Failed to save session"})
		return
	}
}
