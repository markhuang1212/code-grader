package types

type GradeResultStatus int

const (
	GradeResultSuccess GradeResultStatus = iota
	GradeResultWrongAnswer
	GradeResultInternalError
	GradeResultCompilationError
	GradeResultTimeLimitExceed
	GradeResultMemoryExceed
)

type GradeResult struct {
	Status GradeResultStatus
	Msg    string
}
