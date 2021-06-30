package cmd

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/markhuang1212/code-grader/types"
)

func HandleGradeRequest(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
	}

	var gr types.GradeRequest
	body, _ := io.ReadAll(r.Body)

	json.Unmarshal(body, &gr)

}
