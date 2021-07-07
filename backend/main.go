package main

import (
	"github.com/markhuang1212/code-grader/backend/internal/api"
	"github.com/markhuang1212/code-grader/backend/internal/core"
)

func main() {
	cc := core.NewCoreController(3)
	r := api.SetupRouter(cc)
	r.Run(":8080")
}
