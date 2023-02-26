package globals

var Secret = []byte("secret")

const ENV_KEY = "GO_ENV"

const Userkey = "user"

const Username = "username"

const Latest = "latest"

var latestRequestId int = -1

func GetDatabasePath() string {
	/* FIXME: Determine if we want to use an in memory database,
	if it is the case, the database middleware would have to change
	c := os.Getenv(ENV_KEY)
	 if c == "TEST" {
		c = ":memory:"
	} else {
		c = "itu-minitwit.db"
	}
	return c */
	return "itu-minitwit.db"
}

func SetLatestRequestId(requestId int) {
	latestRequestId = requestId
}

func GetLatestRequestId() int {
	return latestRequestId
}
