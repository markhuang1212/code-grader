package cmd

import (
	"errors"
	"os"
)

var (
	TimeLimitExceed   = errors.New("time limit exceeds")
	MemoryLimitExceed = errors.New("memory limit exceeds")
	InternalError     = errors.New("internal error")
	CompilationError  = errors.New("compilation error")
	AppRoot           = os.Getenv("APP_ROOT")
)
