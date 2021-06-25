package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type GradeRequest struct {
	TestCaseId string
	UserCode   string
}

var TestCaseDir string
var TestCaseIds []string

func gradeRequestHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var gr GradeRequest
	err := json.NewDecoder(r.Body).Decode(&gr)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}

func main() {
	TestCaseDir = "/code-grader/testcases"

	files, err := ioutil.ReadDir(TestCaseDir)
	if err != nil {
		return
	}

	TestCaseIds = make([]string, len(files))
	for i, v := range files {
		TestCaseIds[i] = v.Name()
	}

	http.HandleFunc("/grade", gradeRequestHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
