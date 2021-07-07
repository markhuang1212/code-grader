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

func TestExecUserCode1(t *testing.T) {

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

func TestExecUserCode2(t *testing.T) {

	ctx := context.Background()

	tmpDir, err := ioutil.TempDir("/tmp", "")
	assert.Nil(t, err)
	defer os.RemoveAll(tmpDir)

	err = os.Chmod(tmpDir, 0777)
	assert.Nil(t, err)

	gr := types.GradeRequest{
		TestCaseName: "example-1",
		UserCode:     "int main() { return 10; }",
	}

	cr, err := cmd.CompileUserCode(ctx, gr, tmpDir)
	assert.Nil(t, err)
	assert.True(t, cr.Ok)

	er, err := cmd.ExecUserCode(ctx, gr, tmpDir)
	assert.Nil(t, err)
	assert.False(t, er.Ok)
	assert.True(t, er.ExecutionError)

}

func TestExecUserCode3(t *testing.T) {

	ctx := context.Background()

	tmpDir, err := ioutil.TempDir("/tmp", "")
	assert.Nil(t, err)
	defer os.RemoveAll(tmpDir)

	err = os.Chmod(tmpDir, 0777)
	assert.Nil(t, err)

	gr := types.GradeRequest{
		TestCaseName: "example-1",
		UserCode:     "int main() { vector<int> data; while(1) { data.push_back(100); } }",
	}

	cr, err := cmd.CompileUserCode(ctx, gr, tmpDir)
	assert.Nil(t, err)
	assert.True(t, cr.Ok)

	er, err := cmd.ExecUserCode(ctx, gr, tmpDir)
	assert.Nil(t, err)
	assert.False(t, er.Ok)
	assert.True(t, er.MemoryExceed)

}
