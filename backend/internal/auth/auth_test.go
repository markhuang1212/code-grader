package auth_test

import (
	"strings"
	"testing"

	"github.com/markhuang1212/code-grader/backend/internal/auth"
	"github.com/stretchr/testify/assert"
)

var authInfo string = `user1 123456
user2 abcdef
`

func TestAuth(t *testing.T) {
	info := strings.NewReader(authInfo)
	ac, err := auth.NewAuthController(info)
	assert.Nil(t, err)

	assert.True(t, ac.Auth("user1", "123456"))
	assert.True(t, ac.Auth("user2", "abcdef"))

	assert.False(t, ac.Auth("user2", "123456"))
}
