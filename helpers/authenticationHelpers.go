package helpers

import (
	"DevOps/globals"
	"DevOps/model"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) (model.User, error) {
	db := GetTypedDb(c)
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)

	if user != nil {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{"content": "Please logout first"})
		return model.User{}, errors.New("please logout first")
	}

	username := c.PostForm("username")
	password := c.PostForm("password")
	password2 := c.PostForm("password2")
	email := c.PostForm("email")

	if EmptyUserPass(username, password) {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{"content": "You have to enter a value"})
		return model.User{}, errors.New("you have to enter a value")
	}

	if !checkUserPasswords(password, password2) {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{"content": "The two passwords do not match"})
		return model.User{}, errors.New("the two passwords do not match")
	}

	if !checkUserEmail(email) {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{"content": "You have to enter a valid email address"})
		return model.User{}, errors.New("you have to enter a valid email address")
	}

	if checkUsernameExists(db, username) {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{"content": "The username is already taken"})
		return model.User{}, errors.New("the username is already taken")
	}

	pw_hash, err := HashPassword(password)
	if err != nil {
		c.AbortWithStatus(404)
	}
	result, err := db.Exec("insert into user (username, email, pw_hash) values (?, ?, ?)", username, email, pw_hash)

	id, err := result.LastInsertId()

	return model.User{Username: username, UserId: id, Email: email, PwHash: pw_hash}, err
}

func checkUserPasswords(password, password2 string) bool {
	return password == password2
}

func checkUsernameExists(db *sqlx.DB, username string) bool {
	_, err := GetUserId(db, username)
	return err == nil
}

func checkUserEmail(email string) bool {
	return strings.Contains(email, "@")
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
