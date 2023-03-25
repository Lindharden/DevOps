package helpers

import (
	"DevOps/globals"
	model "DevOps/model/gorm"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterUser(db *gorm.DB, username string, password string, password2 string, email string) (model.User, error) {

	if EmptyUserPass(username, password) {
		return model.User{}, errors.New("You have to enter a value")
	}

	if !CheckUserPasswords(password, password2) {
		return model.User{}, errors.New("The two passwords do not match")
	}

	if !CheckUserEmail(email) {
		return model.User{}, errors.New("You have to enter a valid email address")
	}

	if CheckUsernameExists(db, username) {
		return model.User{}, errors.New("The username is already taken")
	}

	pw_hash, err := HashPassword(password)
	if err != nil {
		globals.GetLogger().Errorw("Password hashing failed at user register", "error", err.Error())
		return model.User{}, errors.New("Something went wrong. Minitwit has been notified.")
	}
	user := model.User{Username: username, Email: email, PwHash: pw_hash}
	result := db.Create(&user)

	return user, result.Error
}

func CheckUserPasswords(password, password2 string) bool {
	return password == password2
}

func CheckUsernameExists(db *gorm.DB, username string) bool {
	_, err := GetUserIdGorm(db, username)
	return err == nil
}

func CheckUserEmail(email string) bool {
	return strings.Contains(email, "@")
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}
