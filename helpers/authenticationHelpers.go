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

	if !CheckUserPasswords(password, password2) {
		return model.User{}, errors.New("the two passwords do not match")
	}

	if !CheckUserEmail(email) {
		return model.User{}, errors.New("you have to enter a valid email address")
	}

	if CheckUsernameExists(db, username) {
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

func CheckUserPasswords(password, password2 string) bool {
	return password == password2
}

func CheckUsernameExists(db *sqlx.DB, username string) bool {
	_, err := GetUserId(db, username)
	return err == nil
}

func CheckUserEmail(email string) bool {
	return strings.Contains(email, "@")
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
