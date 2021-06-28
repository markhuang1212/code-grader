// The program read the environment `TEST_CASE_DIR`,
// accepts the user code from stdin, and output the
// compiled executable to stdout.
// In case error occurred, it output the error message
// to stderr, and exit with status 1.
package main

import (
	"os"
)

var testCaseDir string

func main() {
	testCaseDir = os.Getenv("TEST_CASE_DIR")

}
