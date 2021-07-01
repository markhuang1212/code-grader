package cmd

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// Load all existing test case names
func LoadTestcases() ([]string, error) {

	result := []string{}

	dir := filepath.Join(GetAppRoot(), "testcases")
	ret, err := os.ReadDir(dir)

	if err != nil {
		return nil, errors.Wrap(err, "cannot read testcase dir")
	}

	for _, fileEntry := range ret {
		if fileEntry.IsDir() {
			result = append(result, fileEntry.Name())
		}
	}

	return result, nil
}
