package cmd

import (
	"errors"
	"os"
)

var (
	ErrTimeLimitExceed   = errors.New("time limit exceeds")
	ErrMemoryLimitExceed = errors.New("memory limit exceeds")
	ErrInternalError     = errors.New("internal error")
	ErrCompilationError  = errors.New("compilation error")
	AppRoot              = os.Getenv("APP_ROOT")
)
