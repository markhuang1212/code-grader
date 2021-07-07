package types

import "time"

type GradeResultStatus int

const (
	GradeResultSuccess GradeResultStatus = iota
	GradeResultWrongAnswer
	GradeResultExecutionError
	GradeResultInternalError
	GradeResultCompilationError
	GradeResultTimeLimitExceed
	GradeResultMemoryExceed
)

type GradeResult struct {
	Status   GradeResultStatus
	Duration time.Duration
	Msg      string
}
