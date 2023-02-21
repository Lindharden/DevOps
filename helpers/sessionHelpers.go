package helpers

import (
	"DevOps/globals"
	"DevOps/model"
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"errors"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// go binary encoder
func ToGOB64(m model.User) (string, error) {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(m)
	if err != nil {
		fmt.Println(`failed gob Encode`, err)
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b.Bytes()), nil
}

// go binary decoder
func FromGOB64(str string) (model.User, error) {
	m := model.User{}
	by, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println(`failed base64 Decode`, err)
		return m, err
	}
	b := bytes.Buffer{}
	b.Write(by)
	d := gob.NewDecoder(&b)
	err = d.Decode(&m)
	if err != nil {
		fmt.Println(`failed gob Decode`, err)
		return m, err
	}
	return m, nil
}

func GetUserSession(c *gin.Context) (model.User, error) {
	session := sessions.Default(c)
	userSerialized := session.Get(globals.Userkey)
	if userSerialized == nil {
		return model.User{}, errors.New("No session found")
	}
	return FromGOB64(userSerialized.(string))
}

func SetUserSession(c *gin.Context, m model.User) error {
	session := sessions.Default(c)
	userSerialized, err := ToGOB64(m)
	if err != nil {
		return err
	}
	session.Set(globals.Userkey, userSerialized)
	return nil
}
