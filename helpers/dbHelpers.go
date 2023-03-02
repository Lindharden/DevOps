package helpers

import (
	"DevOps/globals"
	model "DevOps/model/gorm"
	"log"

	"github.com/jmoiron/sqlx"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Deprecated: Uses old database
func GetUserId(db *sqlx.DB, username string) (string, error) {
	var id string
	err := db.QueryRow("select user_id from user where username = ?", username).Scan(&id)
	return id, err
}

func GetUserIdGorm(db *gorm.DB, username string) (uint, error) {
	var user model.User
	res := db.Where("username = ?", username).First(&user)
	return user.ID, res.Error
}

// Open connection to DB, and return DB element
func GetGormDbConnection() *gorm.DB {
	dbGorm, err := gorm.Open(sqlite.Open(globals.GetDatabasePath()), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	return dbGorm
}
