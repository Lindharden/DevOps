package test

import (
	"DevOps/routes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type RegisterData struct {
	Email     string
	Username  string
	Password  string
	Password2 string
}

func createRegisterRequest(registerdata RegisterData, router *gin.Engine) *httptest.ResponseRecorder {
	if len(registerdata.Password2) == 0 {
		registerdata.Password2 = registerdata.Password
	}

	if len(registerdata.Email) == 0 {
		registerdata.Email = registerdata.Username + "@mail.com"
	}
	formParams := fmt.Sprintf("username=%s&password=%s&password2=%s&email=%s", registerdata.Username, registerdata.Password, registerdata.Password2, registerdata.Email)
	req, _ := http.NewRequest("POST", "/register",
		strings.NewReader(formParams))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func TestRegisterRoute(t *testing.T) {
	router := routes.SetupRouter()

	r := createRegisterRequest(RegisterData{Username: "user1", Password: "default"}, router)
	assert.Equal(t, http.StatusMovedPermanently, r.Code)

	r = createRegisterRequest(RegisterData{Username: "user1", Password: "default"}, router)
	assert.Equal(t, http.StatusBadRequest, r.Code)
	assert.Contains(t, r.Body.String(), "The username is already taken")

	r = createRegisterRequest(RegisterData{Username: "", Password: "default"}, router)
	assert.Equal(t, http.StatusBadRequest, r.Code)
	assert.Contains(t, r.Body.String(), "You have to enter a value")

	r = createRegisterRequest(RegisterData{Username: "user2", Password: ""}, router)
	assert.Equal(t, http.StatusBadRequest, r.Code)
	assert.Contains(t, r.Body.String(), "You have to enter a value")

	r = createRegisterRequest(RegisterData{Username: "user2", Password: ""}, router)
	assert.Equal(t, http.StatusBadRequest, r.Code)
	assert.Contains(t, r.Body.String(), "You have to enter a value")

	r = createRegisterRequest(RegisterData{Username: "user2", Password: "default", Email: "broken"}, router)
	assert.Equal(t, http.StatusBadRequest, r.Code)
	assert.Contains(t, r.Body.String(), "You have to enter a valid email address")

}

func doLoginRequest(username string, password string, router *gin.Engine) *httptest.ResponseRecorder {
	formParams := fmt.Sprintf("username=%s&password=%s", username, password)
	req, _ := http.NewRequest("POST", "/login",
		strings.NewReader(formParams))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func doLogoutRequest(sessionCookie *http.Cookie, router *gin.Engine) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("GET", "/private/logout", nil)
	req.AddCookie(sessionCookie)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func registerAndLogin(loginData RegisterData, router *gin.Engine) *http.Cookie {
	createRegisterRequest(loginData, router)
	r := doLoginRequest(loginData.Username, loginData.Password, router)
	return r.Result().Cookies()[0]
}

func TestLoginRoute(t *testing.T) {
	router := routes.SetupRouter()

	var user RegisterData = RegisterData{Username: "user1", Password: "default"}
	r := createRegisterRequest(user, router)

	r = doLoginRequest(user.Username, user.Password, router)
	assert.Equal(t, http.StatusMovedPermanently, r.Code)

	sessionCookie := r.Result().Cookies()[0]
	r = doLogoutRequest(sessionCookie, router)
	assert.Equal(t, http.StatusFound, r.Code)

}

func doAddMessageRequest(cookie *http.Cookie, text string, router *gin.Engine) *httptest.ResponseRecorder {
	formParams := fmt.Sprintf("text=%s", text)
	req, _ := http.NewRequest("POST", "/private/message", strings.NewReader(formParams))
	req.AddCookie(cookie)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func getPrivateTimeline(cookie *http.Cookie, router *gin.Engine) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("GET", "/private", nil)
	req.AddCookie(cookie)
	r := httptest.NewRecorder()
	router.ServeHTTP(r, req)
	return r
}

func TestMessageRecording(t *testing.T) {
	router := routes.SetupRouter()
	sessionCookie := registerAndLogin(RegisterData{Username: "user", Password: "bar"}, router)

	doAddMessageRequest(sessionCookie, "test message 1", router)
	doAddMessageRequest(sessionCookie, "test message 2", router)

	req, _ := http.NewRequest("GET", "/private", nil)
	req.AddCookie(sessionCookie)
	r := httptest.NewRecorder()
	router.ServeHTTP(r, req)
	assert.Equal(t, http.StatusOK, r.Code)
	assert.Contains(t, r.Body.String(), "test message 1")
	assert.Contains(t, r.Body.String(), "test message 2")

}

func getPublicTimeline(router *gin.Engine) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("GET", "/public", nil)
	r := httptest.NewRecorder()
	router.ServeHTTP(r, req)
	return r
}

// Action needs to be: follow | unfollow
func followOrUnfollowUser(username string, action string, session *http.Cookie, router *gin.Engine) *httptest.ResponseRecorder {
	uri := fmt.Sprintf("/private/%s/%s", username, action)
	req, _ := http.NewRequest("GET", uri, nil)
	req.AddCookie(session)
	r := httptest.NewRecorder()
	router.ServeHTTP(r, req)
	return r
}

func getUserTimeline(username string, router *gin.Engine) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("GET", fmt.Sprintf("/%s", username), nil)
	r := httptest.NewRecorder()
	router.ServeHTTP(r, req)
	return r
}

func TestTimeLine(t *testing.T) {
	fooUser := RegisterData{Username: "foo1", Password: "default"}
	barUser := RegisterData{Username: "bar", Password: "default"}
	router := routes.SetupRouter()
	fooSessionCookie := registerAndLogin(fooUser, router)
	doAddMessageRequest(fooSessionCookie, "the message by foo", router)

	doLogoutRequest(fooSessionCookie, router)

	barSessionCookie := registerAndLogin(barUser, router)
	doAddMessageRequest(barSessionCookie, "the message by bar", router)

	//check that public timeline contains both messages
	r := getPublicTimeline(router)
	assert.Contains(t, r.Body.String(), "the message by foo")
	assert.Contains(t, r.Body.String(), "the message by bar")

	//bars timeline should contain bars message but not foos
	r = getPrivateTimeline(barSessionCookie, router)
	assert.NotContains(t, r.Body.String(), "the message by foo")
	assert.Contains(t, r.Body.String(), "the message by bar")

	//follow foo
	followOrUnfollowUser(fooUser.Username, "follow", barSessionCookie, router)

	//check that both messages are now on the pr
	r = getPrivateTimeline(barSessionCookie, router)
	assert.Contains(t, r.Body.String(), "the message by foo")
	assert.Contains(t, r.Body.String(), "the message by bar")

	//bar user timeline only shows bar message
	r = getUserTimeline(barUser.Username, router)
	assert.NotContains(t, r.Body.String(), "the message by foo")
	assert.Contains(t, r.Body.String(), "the message by bar")

	//foo user timeline only shows foo message
	r = getUserTimeline(fooUser.Username, router)
	assert.Contains(t, r.Body.String(), "the message by foo")
	assert.NotContains(t, r.Body.String(), "the message by bar")

	//check that unfollow worked
	r = followOrUnfollowUser(fooUser.Username, "unfollow", barSessionCookie, router)
	r = getPrivateTimeline(barSessionCookie, router)
	assert.NotContains(t, r.Body.String(), "the message by foo")
	assert.Contains(t, r.Body.String(), "the message by bar")
}
