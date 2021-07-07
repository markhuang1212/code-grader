package grader

import (
	"log"
	"os"
	"path/filepath"
	"sort"

	"github.com/markhuang1212/code-grader/backend/internal/util"
)

var testcases []string

// Load all existing test case names
func LoadTestcases() []string {

	if testcases == nil {

		dir := filepath.Join(util.GetAppRoot(), "testcases")
		ret, err := os.ReadDir(dir)

		if err != nil {
			log.Fatal("cannot read testcases")
		}

		for _, fileEntry := range ret {
			if fileEntry.IsDir() {
				testcases = append(testcases, fileEntry.Name())
			}
		}

		sort.Strings(testcases)

	}

	return testcases
}

func IsTestcase(name string) bool {

	if testcases == nil {
		LoadTestcases()
	}

	idx := sort.SearchStrings(testcases, name)

	if idx == len(testcases) || testcases[idx] != name {
		return false
	}

	return true

}
