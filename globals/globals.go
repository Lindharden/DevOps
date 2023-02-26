package globals

var Secret = []byte("secret")

const Userkey = "user"

const Username = "username"

const Latest = "latest"

const DATABASE = "itu-minitwit.db"

var latestRequestId int = -1

func SetLatestRequestId(requestId int) {
	latestRequestId = requestId
}

func GetLatestRequestId() int {
	return latestRequestId
}
