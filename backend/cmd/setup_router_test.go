package cmd_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/markhuang1212/code-grader/backend/cmd"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := cmd.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
