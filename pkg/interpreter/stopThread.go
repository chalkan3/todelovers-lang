package interpreter

import (
	"fmt"
	"mary_guica/pkg/tvm"
)

type stopThread struct {
	*base
}

func NewStopThread(vm *tvm.TVM) Command {
	return &thread{
		&base{
			tvm: vm,
		},
	}
}

func (c *stopThread) Execute(instruction byte, threadID int, args ...interface{}) {
	fmt.Println("Thread is Over")
}
