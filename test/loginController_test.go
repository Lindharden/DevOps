package test

import (
	"DevOps/routes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createRegisterRequest() (*http.Request, error) {
	return http.NewRequest("POST", "/register",
		strings.NewReader("username=asd&password=123&password2=123&email=test@mail"))
}

func TestRegisterRoute(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := createRegisterRequest()
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	router.ServeHTTP(w, req)

	assert.Equal(t, 301, w.Code)
}
