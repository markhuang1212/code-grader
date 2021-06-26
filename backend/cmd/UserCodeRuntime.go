package cmd

type UserCodeExecutionResult struct {
	Status ExecutionStatus
	Msg    string
}

type ExecutionStatus int

const (
	ExecutionStatusSuccess ExecutionStatus = 0
	ExecutionStatusCompilationError
	ExecutionStatusWrongResult
	ExecutionStatusMemoryLimitExceed
	ExecutionStatusTimeLimitExceed
)

func EvaluateUserCode(gq GradeRequest) UserCodeExecutionResult {
	result := UserCodeExecutionResult{
		Status: 0,
		Msg:    "",
	}

	return result
}
