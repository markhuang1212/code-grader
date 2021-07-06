package cmd_test

import (
	"context"
	"io/ioutil"
	"os"
	"testing"

	"github.com/markhuang1212/code-grader/backend/cmd"
	"github.com/markhuang1212/code-grader/backend/types"
	"github.com/stretchr/testify/assert"
)

func TestExecUserCode(t *testing.T) {

	ctx := context.Background()

	tmpDir, err := ioutil.TempDir("/tmp", "")
	assert.Nil(t, err)
	defer os.RemoveAll(tmpDir)

	err = os.Chmod(tmpDir, 0777)
	assert.Nil(t, err)

	gr := types.GradeRequest{
		TestCaseName: "example-1",
		UserCode:     "int main() { cout << \"Hello\" << endl; }",
	}

	cr, err := cmd.CompileUserCode(ctx, gr, tmpDir)
	assert.Nil(t, err)
	assert.True(t, cr.Ok)

	er, err := cmd.ExecUserCode(ctx, gr, tmpDir)
	assert.Nil(t, err)
	assert.True(t, er.Ok)

}
