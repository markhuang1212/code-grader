package types

type GradeResultStatus int

const (
	GradeResultSuccess GradeResultStatus = iota
	GradeResultWrongAnswer
	GradeResultCompilationError
	GradeResultTimeLimitExceed
	GradeResultMemoryExceed
)

type GradeResult struct {
	Status GradeResultStatus
	Msg    string
}
