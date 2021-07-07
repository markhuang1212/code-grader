package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/markhuang1212/code-grader/backend/internal/api"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := api.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
