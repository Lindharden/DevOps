package globals

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var Secret = []byte("secret")

const ENV_KEY = "GO_ENV"

const Userkey = "user"

const Username = "username"

const Latest = "latest"

var latestRequestId int = -1

var db *sqlx.DB

func GetDatabase() *sqlx.DB {
	return db
}

func SetDatabase(database *sqlx.DB) {
	db = database
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
