package test

import (
	"DevOps/routes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type RegisterData struct {
	Email     string
	Username  string
	Password  string
	Password2 string
}

func createRegisterRequest(registerdata RegisterData) (*http.Request, error) {
	if len(registerdata.Password2) == 0 {
		registerdata.Password2 = registerdata.Password
	}

	return http.NewRequest("POST", "/register",
		strings.NewReader("username=asd&password=123&password2=123&email=test@mail"))
}

func TestRegisterRoute(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := createRegisterRequest(RegisterData{})
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	router.ServeHTTP(w, req)

	assert.Equal(t, 301, w.Code)
}
