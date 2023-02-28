package helpers

import (
	"DevOps/globals"
	"log"

	"github.com/jmoiron/sqlx"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Convenience method to look up the id for a username.
func GetUserId(db *sqlx.DB, username string) (string, error) {
	var id string
	err := db.QueryRow("select user_id from user where username = ?", username).Scan(&id)
	return id, err
}

// Open connection to DB, and return DB element
func GetGormDbConnection() *gorm.DB {
	dbGorm, err := gorm.Open(sqlite.Open(globals.GetDatabasePath()), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	return dbGorm
}
