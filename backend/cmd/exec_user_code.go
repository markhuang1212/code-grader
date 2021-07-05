package cmd

import (
	"context"

	"github.com/markhuang1212/code-grader/backend/types"
)

const imageExec = "markhuang1212/code-grader/runtime-exec:latest"

type ExecUserCodeResult struct {
	Ok  bool
	Msg string
}

func ExecUserCode(ctx context.Context, gr types.GradeRequest) (*ExecUserCodeResult, error) {

	result := &ExecUserCodeResult

	return result, nil

}
