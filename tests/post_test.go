package tests

import (
	"github.com/stretchr/testify/assert"
	"go_code/gintest/routes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPosts(t *testing.T) {
	// 注意：在单元测试时，如果需要用到配置信息，需要在测试时引入配置，比如
	// go test tests/main_test.go tests/post_test.go -v
	r := routes.RegisterRouters()
	body := `{
		"page": 1,
		"page_size": 20,
		"order_key": "id",
		"order_direction": "desc",
		"community_id": "0"
	}`

	// url 必须写全，`/api/v1/posts` 而不是 `api/v1/posts`
	request, _ := http.NewRequest(http.MethodGet, "/api/v1/posts", strings.NewReader(body))
	w := httptest.NewRecorder()

	r.ServeHTTP(w, request)
	// 可以通过 Contains 判断返回的状态码
	assert.Contains(t, w.Body.String(), "40100")

	// 或者
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyNDYyNzg5Mzk4ODQ5NTM2MCwidXNlcm5hbWUiOiJwaHBlciIsImlzcyI6IumahuWGrOW8uiIsImV4cCI6MTY3MTMyMjIwMn0.EJMrNywiKozwKwAOt8OjTtAkaNUMujt_NmikVMlJsxU"
	request.Header.Add("Authorization", "Bearer " + token)
	r.ServeHTTP(w, request)
	t.Log("结果：", w.Body.String())
	assert.Contains(t, w.Body.String(), "200")
}
