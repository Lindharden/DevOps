package globals

import (
	"DevOps/helpers"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var Secret = []byte("secret")

const Userkey = "user"

const Username = "username"

const Latest = "latest"

var latestReuqestId int = -1;

func SaveRequest(c *gin.Context) {
	session := sessions.Default(c)
	reqJson, _ := json.Marshal(c.Request)
	session.Set(Latest, reqJson)
	if err := session.Save(); err != nil {
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{"content": "Failed to save session"})
		return
	}
}

func SetLatestRequestId(requestId int){
	latestReuqestId = requestId
}

func GetLatestRequestId() int {
	return latestReuqestId
}

// get user id from a username
func GetUserId(username string, c *gin.Context) int {
	db := helpers.GetTypedDb(c)

	user := struct{ UserId int }{} // save user_id here
	query := fmt.Sprintf("SELECT user.user_id FROM user WHERE username = %s", username)
	db.Select(&user, query, 1)

	return user.UserId
}
