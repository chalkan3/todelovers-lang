package interpreter

import (
	"mary_guica/pkg/tvm"
	"os"
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

func (c *halt) Execute(instruction byte) {
	os.Exit(0)
}
