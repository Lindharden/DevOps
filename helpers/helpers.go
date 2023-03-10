package helpers

import (
	model "DevOps/model"
	gormModel "DevOps/model/gorm"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

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

func ValidatePassword(db *gorm.DB, username, password string) bool {
	var user gormModel.User
	res := db.Where(gormModel.User{Username: username}).First(&user)
	if res.Error != nil {

		return false
	}

	return CheckPasswordHash(password, user.PwHash)
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

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}
