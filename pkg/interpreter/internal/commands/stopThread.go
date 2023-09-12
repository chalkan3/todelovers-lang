package commands

import (
	"mary_guica/pkg/tvm"
)

type stopThread struct {
	*base
}

func NewStopThread(vm *tvm.TVM) Command {
	return &stopThread{
		&base{
			tvm: vm,
		},
	}
}

func (c *stopThread) setParentDont(threadID int) {
	current := c.tvm.GetThreadManager().GetThread(threadID)
	if current.ParentID() != -1 {
		parent := c.tvm.GetThreadManager().GetParent(current.GetID())
		if parent.Waiting() {
			parent.SetWaitRelease()
			c.setParentDont(current.ParentID())

		}

	}

}

func (c *stopThread) Execute(instruction byte, threadID int, args ...interface{}) {
	c.setParentDont(threadID)
	c.tvm.GetThreadManager().GetThread(threadID).Done()

}
