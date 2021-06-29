// The program read the environment `TEST_CASE_DIR`,
// accepts the tompiled executable from stdin
// If the test case passes, it exit with status 0
// In case error occurred, it output the error message
// to stderr, and exit with status 1.

package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/markhuang1212/code-grader/types"
)

func main() {

	prog, err := os.OpenFile("/tmp/a.out", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	io.Copy(prog, os.Stdin)
	prog.Close()

	testCaseDir := os.Getenv("TEST_CASE_DIR")
	testCaseJson, err := os.ReadFile(filepath.Join(testCaseDir, "testcase.json"))
	if err != nil {
		panic(err)
	}

	var testCaseOpt types.TestCaseOptions
	json.Unmarshal(testCaseJson, &testCaseOpt)

	inputFile, err := os.Open(filepath.Join(testCaseDir, testCaseOpt.RuntimeOptions.StdinPath))
	if err != nil {
		panic(err)
	}

	answerFile, err := os.Open(filepath.Join(testCaseDir, testCaseOpt.RuntimeOptions.StdoutPath))
	if err != nil {
		panic(err)
	}
	answerScanner := bufio.NewScanner(answerFile)
	answerScanner.Split(bufio.ScanWords)

	cmd := exec.Command("/tmp/a.out")

	progOutput, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	progOutputScanner := bufio.NewScanner(progOutput)
	progOutputScanner.Split(bufio.ScanWords)

	cmd.Stdin = inputFile
	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	for {

		r1 := answerScanner.Scan()
		r2 := progOutputScanner.Scan()

		if r1 == false && r2 == false {
			os.Exit(0)
		}

		if r1 == false || r2 == false {
			log.Println("Wrong Answer!")
			os.Exit(1)
		}

		w1 := answerScanner.Text()
		w2 := progOutputScanner.Text()

		if w1 != w2 {
			log.Println("Wrong Answer!")
			os.Exit(1)
		}

	}

}
