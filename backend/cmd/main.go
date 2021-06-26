package main

import (
	"encoding/json"
	"fmt"
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

	switch r.Method {

	case "GET":
		fmt.Fprintf(w, "Usage: POST { TestCaseId: string, UserCode: string }")

	case "POST":

		var gr GradeRequest
		err := json.NewDecoder(r.Body).Decode(&gr)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if gr.TestCaseId == "" || gr.UserCode == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	default:
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
