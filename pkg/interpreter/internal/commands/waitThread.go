package commands

import (
	"mary_guica/pkg/tvm"
)

type waitThread struct {
	*base
}

func NewWaitThread(vm *tvm.TVM) Command {
	return &waitThread{
		&base{
			tvm: vm,
		},
	}
}

func (c *waitThread) setParentWait(threadID int) {
	current := c.GetCurrentThread(threadID)
	current.SetWait()

	if current.ParentID() != -1 {
		c.setParentWait(current.ParentID())
		return

	}

}

func (c *waitThread) Execute(instruction byte, threadID int, args ...interface{}) {
	c.setParentWait(threadID)
}
