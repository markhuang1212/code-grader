package cmd_test

import (
	"context"
	"testing"

	"github.com/markhuang1212/code-grader/backend/cmd"
	"github.com/markhuang1212/code-grader/backend/types"
	"github.com/stretchr/testify/assert"
)

func TestGradeUserCode1(t *testing.T) {
	ctx := context.Background()
	gr := types.GradeRequest{
		TestCaseName: "example-1",
		UserCode:     "int main() { cout << \"Hello\" << endl; }",
	}
	result, err := cmd.GradeUserCode(ctx, gr)
	assert.Nil(t, err)
	assert.Equal(t, result.Status, types.GradeResultSuccess)
}

func TestGradeUserCode2(t *testing.T) {
	ctx := context.Background()
	gr := types.GradeRequest{
		TestCaseName: "example-1",
		UserCode:     "int main() { while (1) { } }",
	}
	result, err := cmd.GradeUserCode(ctx, gr)
	assert.Nil(t, err)
	assert.Equal(t, result.Status, types.GradeResultTimeLimitExceed)
}
