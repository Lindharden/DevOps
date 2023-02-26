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
