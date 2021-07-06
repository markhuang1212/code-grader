package cmd

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/markhuang1212/code-grader/backend/types"
	"github.com/pkg/errors"
)

// this function wraps CompileUserCode and GradeUserCode
func GradeUserCode(ctx context.Context, gr types.GradeRequest) (*types.GradeResult, error) {

	result := types.GradeResult{}

	tmpDir, err := ioutil.TempDir("/tmp", "cdgr_")
	if err != nil {
		return nil, errors.Wrap(err, "cannot create tmpDir")
	}
	defer os.RemoveAll(tmpDir)
	os.Chmod(tmpDir, 0777)

	cr, err := CompileUserCode(ctx, gr, tmpDir)
	if err != nil {
		return nil, errors.Wrap(err, "cannot compile")
	}

	if !cr.Ok {
		result.Status = types.GradeResultCompilationError
		result.Msg = cr.Msg
		return &result, nil
	}

	er, err := ExecUserCode(ctx, gr, tmpDir)
	if err != nil {
		return nil, errors.Wrap(err, "cannot execute")
	}

	if !er.Ok {
		result.Msg = er.Msg
		if er.MemoryExceed {
			result.Status = types.GradeResultMemoryExceed
		} else if er.TimeExceed {
			result.Status = types.GradeResultTimeLimitExceed
		} else {
			result.Status = types.GradeResultWrongAnswer
		}
		return &result, nil
	}

	result.Status = types.GradeResultSuccess
	result.Msg = "success"
	return &result, nil

}
