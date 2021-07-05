package cmd_test

import (
	"testing"

	"github.com/markhuang1212/code-grader/backend/cmd"
	"github.com/stretchr/testify/assert"
)

func TestLoadTestcases(t *testing.T) {
	ret := cmd.LoadTestcases()
	for _, v := range ret {
		if v == "example-1" {
			return
		}
	}
	t.Errorf("wrong result")
}

func TestIsTestcase(t *testing.T) {
	ret := cmd.IsTestcase("example-1")
	assert.True(t, ret)
}
