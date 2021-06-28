// The program read the environment `TEST_CASE_DIR`,
// accepts the user code from stdin, and output the
// compiled executable to stdout.
// In case error occurred, it output the error message
// to stderr, and exit with status 1.
package main

import (
	"encoding/json"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/markhuang1212/code-grader/types"
)

var testCaseDir string

func main() {
	testCaseDir = os.Getenv("TEST_CASE_DIR")

	testCaseJson, err := os.ReadFile(filepath.Join(testCaseDir, "testcase.json"))
	if err != nil {
		panic(err)
	}
	var testCaseOptions types.TestCaseOptions
	json.Unmarshal(testCaseJson, &testCaseOptions)

	// compileCmd := exec.Command("g++", testCaseOptions.CompilerOptions.Flags...)

}
