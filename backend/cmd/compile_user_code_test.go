package cmd_test

import (
	"context"
	"testing"

	"github.com/markhuang1212/code-grader/backend/cmd"
	"github.com/markhuang1212/code-grader/backend/types"
	"github.com/stretchr/testify/assert"
)

func TestCompileUserCode(t *testing.T) {
	ctx := context.Background()

	gr1 := types.GradeRequest{
		TestCaseName: "example-1",
		UserCode:     " ",
	}

	result, err := cmd.CompileUserCode(ctx, gr1)
	assert.Nil(t, err)
	assert.False(t, result.Ok)
	t.Log(result)

	gr2 := types.GradeRequest{
		TestCaseName: "example-1",
		UserCode:     "int main() { cout << \"Hello\" << endl; }",
	}

	result, err = cmd.CompileUserCode(ctx, gr2)
	assert.Nil(t, err)
	assert.True(t, result.Ok)
	t.Log(result)

	gr3 := types.GradeRequest{
		TestCaseName: "some-random-name",
		UserCode:     " ",
	}
	_, err = cmd.CompileUserCode(ctx, gr3)
	assert.ErrorIs(t, err, types.ErrNoTestCase)

}
