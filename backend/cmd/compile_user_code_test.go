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

	out, err := cmd.CompileUserCode(ctx, gr1)
	assert.ErrorIs(t, err, cmd.ErrCompilationError)
	t.Log(out)

	gr2 := types.GradeRequest{
		TestCaseName: "example-1",
		UserCode:     "int main() { cout << \"Hello\" << endl; }",
	}

	out, err = cmd.CompileUserCode(ctx, gr2)
	assert.Equal(t, err, nil)
	t.Log(out)

}
