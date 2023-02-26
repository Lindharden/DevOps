package helpers

import (
	"DevOps/model"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CheckUserPass(username, password string) bool {
	return true
}

func EmptyUserPass(username, password string) bool {
	return strings.Trim(username, " ") == "" || strings.Trim(password, " ") == ""
}

func CheckUserPasswords(password, password2 string) bool {
	return password == password2
}

func CheckUsernameExists(db *sqlx.DB, username string) bool {
	_, err := GetUserId(db, username)
	return err == nil
}

func ValidatePassword(db *sqlx.DB, username, password string) bool {
	var pw_hash string
	err := db.QueryRow("select pw_hash from user where username = ?", username).Scan(&pw_hash)
	if err != nil {
		return false
	}

	return CheckPasswordHash(password, pw_hash)
}

func CheckUserEmail(email string) bool {
	return strings.Contains(email, "@")
}

func GetTypedDb(c *gin.Context) *sqlx.DB {
	db := c.MustGet("db").(*sqlx.DB)
	return db
}

// Return the gravatar image for the given email address.
func GravatarUrl(email string) string {
	size := 48
	return fmt.Sprintf("http://www.gravatar.com/avatar/%s?d=identicon&s=%d", getMD5Hash(strings.ToLower(strings.TrimSpace(email))), size)
}

// Returns argument as a MD5 hashed string
func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func FormatDatetime(unixTime int64) string {
	t := time.Unix(unixTime, 0)
	return t.Format("2 Jan 2006 15:04")
}

func RequestedUserExists(requestedUser string, users []model.User) bool {
	return len(users) >= 0 && len(requestedUser) != 0
}
