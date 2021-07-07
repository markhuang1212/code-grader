package api_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/markhuang1212/code-grader/backend/internal/api"
	"github.com/markhuang1212/code-grader/backend/internal/core"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := api.SetupRouter(core.NewCoreController(1))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestGrading(t *testing.T) {
	router := api.SetupRouter(core.NewCoreController(1))

	w := httptest.NewRecorder()

	body := `
		{
			"TestCaseName": "example-1",
			"UserCode": "int main() { return 0; }"
		}
	`
	req, _ := http.NewRequest("POST", "/api/v1/grade", strings.NewReader(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusAccepted, w.Code)
	loc := w.Header().Get("Location")
	assert.NotEmpty(t, loc)
}
