package helpers

import (
	"DevOps/model"
	"errors"
	"strings"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(db *sqlx.DB, username string, password string, password2 string, email string) (model.User, error) {

	if EmptyUserPass(username, password) {
		return model.User{}, errors.New("you have to enter a value")
	}

	if !checkUserPasswords(password, password2) {
		return model.User{}, errors.New("the two passwords do not match")
	}

	if !checkUserEmail(email) {
		return model.User{}, errors.New("you have to enter a valid email address")
	}

	if checkUsernameExists(db, username) {
		return model.User{}, errors.New("the username is already taken")
	}

	pw_hash, err := HashPassword(password)
	if err != nil {
		panic("password hashing failed")
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
