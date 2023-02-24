package helpers

import (
	"DevOps/globals"
	"DevOps/model"
	"encoding/json"
	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetUserSession(c *gin.Context) (model.User, error) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)
	if user != nil {
		var deserialized model.User
		err := json.Unmarshal(user.([]byte), &deserialized)
		return deserialized, err
	}
	return model.User{}, errors.New("No session active")
}

func SetUserSession(c *gin.Context, m model.User) error {
	session := sessions.Default(c)
	serialized, err := json.Marshal(m)
	if err != nil {
		return err
	}
	session.Set(globals.Userkey, serialized)
	return session.Save()
}

func TerminateUserSession(c *gin.Context) error {
	session := sessions.Default(c)
	session.Delete(globals.Userkey)
	session.Clear()
	return session.Save()
}
