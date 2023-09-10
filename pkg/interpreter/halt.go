package interpreter

import (
	"mary_guica/pkg/tvm"
)

type halt struct {
	*base
}

func NewHalt(vm *tvm.TVM) Command {
	return &halt{
		&base{
			tvm: vm,
		},
	}
}

func (c *halt) Execute(instruction byte, threadID int, args ...interface{}) {
	c.GetCurrentThread(threadID).SetDone()
}
