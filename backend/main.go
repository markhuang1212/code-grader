package main

import (
	"github.com/markhuang1212/code-grader/backend/cmd"
)

func main() {
	r := cmd.SetupRouter()
	r.Run(":8080")
}
