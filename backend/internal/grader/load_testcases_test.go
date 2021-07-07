package grader_test

import (
	"testing"

	"github.com/markhuang1212/code-grader/backend/internal/grader"
	"github.com/stretchr/testify/assert"
)

func TestLoadTestcases(t *testing.T) {
	ret := grader.LoadTestcases()
	for _, v := range ret {
		if v == "example-1" {
			return
		}
	}
	t.Errorf("wrong result")
}

func TestIsTestcase(t *testing.T) {
	ret := grader.IsTestcase("example-1")
	assert.True(t, ret)
}
