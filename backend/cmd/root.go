package cmd

import (
	"io/ioutil"
	"net/http"
)

var TestCaseDir string
var TestCaseIds []string

func Execute() error {

	TestCaseDir = "/code-grader/testcases"

	files, err := ioutil.ReadDir(TestCaseDir)
	if err != nil {
		return err
	}

	TestCaseIds = make([]string, len(files))
	for i, v := range files {
		TestCaseIds[i] = v.Name()
	}

	http.HandleFunc("/grade", GradeRequestHandler)
	err = http.ListenAndServe(":8080", nil)

	return err
}
