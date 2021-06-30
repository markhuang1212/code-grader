package cmd_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/markhuang1212/code-grader/backend/cmd"
	"github.com/markhuang1212/code-grader/types"
)

func TestCompileUserCode(t *testing.T) {
	ctx := context.Background()
	gr := types.GradeRequest{
		TestCaseName: "example-1",
		UserCode:     "",
	}

	_, err := cmd.CompileUserCode(ctx, gr)
	fmt.Println(err)
}
