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

func TestCompileUserCode1(t *testing.T) {
	ctx := context.Background()

	tmpDir, err := ioutil.TempDir("/tmp", "")
	assert.Nil(t, err)
	defer os.RemoveAll(tmpDir)

	err = os.Chmod(tmpDir, 0777)
	assert.Nil(t, err)

	gr1 := types.GradeRequest{
		TestCaseName: "example-1",
		UserCode:     " ",
	}

	result, err := cmd.CompileUserCode(ctx, gr1, tmpDir)
	assert.Nil(t, err)
	assert.False(t, result.Ok)
	t.Log(result)

}

func TestCompileUserCode2(t *testing.T) {

	ctx := context.Background()

	tmpDir, err := ioutil.TempDir("/tmp", "")
	assert.Nil(t, err)
	defer os.RemoveAll(tmpDir)

	err = os.Chmod(tmpDir, 0777)
	assert.Nil(t, err)

	gr2 := types.GradeRequest{
		TestCaseName: "example-1",
		UserCode:     "int main() { cout << \"Hello\" << endl; }",
	}

	result, err := cmd.CompileUserCode(ctx, gr2, tmpDir)
	assert.Nil(t, err)
	assert.True(t, result.Ok)
	t.Log(result)

}
