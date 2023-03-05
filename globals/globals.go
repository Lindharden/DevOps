package globals

import (
	model "DevOps/model/gorm"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
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

func GetDatabasePath() gorm.Dialector {
	if os.Getenv(ENV_KEY) == "production" {
		connectionString := fmt.Sprintf("postgresql://%s:%s@postgres/db",
			os.Getenv("POSTGRES_USERNAME"),
			os.Getenv("POSTGRES_PASSWORD"))
		fmt.Println(connectionString)
		return postgres.Open(connectionString)
	}
	return sqlite.Open("itu-minitwit.db")
}

func SetLatestRequestId(requestId int) {
	latestRequestId = requestId
}

func GetLatestRequestId() int {
	return latestRequestId
}
