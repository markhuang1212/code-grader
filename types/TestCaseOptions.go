package types

type TestCaseOptions struct {
	TestCaseName      string
	PreprocessOptions PreprocessOptions
	CompilerOptions   CompilerOptions
	RuntimeOptions    RuntimeOptions
}

type PreprocessOptions struct {
	AppendCodePath  string
	PrependCodePath string
}

type CompilerOptions struct {
	Flags []string
}

type RuntimeOptions struct {
	StdinPath    string
	StdoutPath   string
	MemoryLimit  int
	RuntimeLimit int
}
