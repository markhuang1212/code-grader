package cmd_test

import (
	"context"
	"testing"

	"github.com/markhuang1212/code-grader/backend/cmd"
	"github.com/markhuang1212/code-grader/backend/types"
	"github.com/stretchr/testify/assert"
)

// success
func TestGradeUserCode1(t *testing.T) {
	ctx := context.Background()
	gr := types.GradeRequest{
		TestCaseName: "example-1",
		UserCode:     "int main() { cout << \"Hello\" << endl; }",
	}
	result, err := cmd.GradeUserCode(ctx, gr)
	assert.Nil(t, err)
	assert.Equal(t, types.GradeResultSuccess, result.Status)
}

// wrong answer
func TestGradeUserCode2(t *testing.T) {
	ctx := context.Background()
	gr := types.GradeRequest{
		TestCaseName: "example-1",
		UserCode:     "int main() { cout << \"Helloooo\" << endl; }",
	}
	result, err := cmd.GradeUserCode(ctx, gr)
	assert.Nil(t, err)
	assert.Equal(t, types.GradeResultWrongAnswer, result.Status)
}

// time limit exceed
func TestGradeUserCode3(t *testing.T) {
	ctx := context.Background()
	gr := types.GradeRequest{
		TestCaseName: "example-1",
		UserCode:     "int main() { while (1) { } }",
	}
	result, err := cmd.GradeUserCode(ctx, gr)
	assert.Nil(t, err)
	assert.Equal(t, types.GradeResultTimeLimitExceed, result.Status)
}

// memory limit exceed
func TestGradeUserCode4(t *testing.T) {
	ctx := context.Background()
	gr := types.GradeRequest{
		TestCaseName: "example-1",
		UserCode:     "int main() { vector<int> data; while (1) { data.push_back(1024); } }",
	}
	result, err := cmd.GradeUserCode(ctx, gr)
	assert.Nil(t, err)
	assert.Equal(t, types.GradeResultMemoryExceed, result.Status)
}
