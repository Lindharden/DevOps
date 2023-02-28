package globals

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
)

var Secret = []byte("secret")

const ENV_KEY = "GO_ENV"

const Userkey = "user"

const Username = "username"

const Latest = "latest"

var latestRequestId int = -1

var db *sqlx.DB

var gormDb *gorm.DB

func GetDatabase() *sqlx.DB {
	return db
}

func GetGormDatabase() *gorm.DB {
	return gormDb
}
func SetDatabase(database_old *sqlx.DB, database *gorm.DB) {
	db = database_old
	gormDb = database
}

func GetDatabasePath() string {
	return "itu-minitwit.db"
}

func SetLatestRequestId(requestId int) {
	latestRequestId = requestId
}

func GetLatestRequestId() int {
	return latestRequestId
}
