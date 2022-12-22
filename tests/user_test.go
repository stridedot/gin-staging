package tests

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go_code/gintest/app/controllers/user"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testRegister(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	url := "/api/v1/register"
	r.POST(url, user.SignUp)

	body := `{
		"username": "test",
		"password": "test123"
	}`

	request, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)
	t.Log(w.Body)
	assert.Equal(t, 200, w.Code)
}

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	url := "/api/v1/login"
	r.POST(url, user.SignIn)

	body := `{
		"username": "test",
		"password": "test123"
	}`

	request, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)
	t.Log(w.Body)
	assert.Equal(t, 200, w.Code)
}
