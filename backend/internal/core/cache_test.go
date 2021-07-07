package core_test

import (
	"testing"

	"github.com/markhuang1212/code-grader/backend/internal/core"
	"github.com/markhuang1212/code-grader/backend/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestCache(t *testing.T) {
	c := core.NewGradeResultCache()
	c.Add("id1", types.GradeResult{
		Status: types.GradeResultCompilationError,
		Msg:    "Hello",
	})
	ret, ok := c.Get("id1")
	assert.True(t, ok)
	assert.Equal(t, "Hello", ret.Msg)
}
