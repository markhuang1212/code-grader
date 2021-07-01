package cmd

import (
	"path"
	"runtime"
)

func GetAppRoot() string {
	_, filename, _, _ := runtime.Caller(0)
	return path.Join(path.Dir(filename), "../../")
}
