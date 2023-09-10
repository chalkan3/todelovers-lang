package tvm

type Interpreter interface {
	Handle(instruction byte)
}
