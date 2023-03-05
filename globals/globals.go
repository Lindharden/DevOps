package globals

import (
	model "DevOps/model/gorm"

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
func SetDatabase(database *gorm.DB) {

	gormDb = database
	gormDb.AutoMigrate(&model.User{})
	gormDb.AutoMigrate(&model.Message{})
	gormDb.AutoMigrate(&model.Following{})
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
