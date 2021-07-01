package types

type TestCaseOptions struct {
	RuntimeOptions RuntimeOptions
}

type RuntimeOptions struct {
	MemoryLimit  int
	RuntimeLimit int
}
