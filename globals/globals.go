package globals

import "net/http"

var Secret = []byte("secret")

const Userkey = "user"

var LatestRequest *http.Request = nil
