package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GradeRequest struct {
	TestCaseId string
	UserCode   string
}

func GradeRequestHandler(w http.ResponseWriter, r *http.Request) {

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
