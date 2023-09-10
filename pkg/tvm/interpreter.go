package tvm

type Interpreter interface {
	Handle(instruction byte, threadID int, args ...interface{})
}
