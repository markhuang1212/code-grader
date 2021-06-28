package types

type GradeResult int

const (
	GradeResultSuccess GradeResult = iota
	GradeResultWrongAnswer
	GradeResultCompilationError
	GradeResultTimeLimitExceed
	GradeResultMemoryExceed
)

type GradeResultMsg struct {
	Result GradeResult
	Msg    string
}
