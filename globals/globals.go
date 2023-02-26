package globals

import "github.com/jmoiron/sqlx"

var Secret = []byte("secret")

const ENV_KEY = "GO_ENV"

const Userkey = "user"

const Username = "username"

const Latest = "latest"

var latestRequestId int = -1

var DB *sqlx.DB

func GetDatabasePath() string {
	return "itu-minitwit.db"
}

func SetLatestRequestId(requestId int) {
	latestRequestId = requestId
}

func GetLatestRequestId() int {
	return latestRequestId
}
