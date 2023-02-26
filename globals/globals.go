package globals

import (
	"os"
)

var Secret = []byte("secret")

const Userkey = "user"

const Username = "username"

const Latest = "latest"

var latestRequestId int = -1

func GetDatabasePath() string {
	c := os.Getenv("GO_ENV")
	if c == "TEST" {
		c = ":memory:"
	} else {
		c = "itu-minitwit.db"
	}
	return c
}

func SetLatestRequestId(requestId int) {
	latestRequestId = requestId
}

func GetLatestRequestId() int {
	return latestRequestId
}
