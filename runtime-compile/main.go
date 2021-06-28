// The program read the environment `TEST_CASE_DIR`,
// accepts the user code from stdin, and output the
// compiled executable to stdout.
// In case error occurred, it output the error message
// to stderr, and exit with status 1.
package main

import (
	"encoding/json"
	"errors"
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

	args := append(testCaseOptions.CompilerOptions.Flags, "-o-", "-x", "c++", "-")
	compileCmd := exec.Command("g++", args...)

	compileCmdStdin, err := compileCmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	compileCmdStdout, err := compileCmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	compileCmdStderr, err := compileCmd.StderrPipe()
	if err != nil {
		panic(err)
	}

	go func() {
		io.Copy(compileCmdStdin, os.Stdin)
		compileCmdStdin.Close()
	}()

	go io.Copy(os.Stderr, compileCmdStderr)
	go io.Copy(os.Stdout, compileCmdStdout)

	err = compileCmd.Run()

	var exitError *exec.ExitError

	if errors.As(err, &exitError) {
		os.Exit(exitError.ExitCode())
	}

}
