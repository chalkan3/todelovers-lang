package interpreter

import (
	"mary_guica/pkg/tvm"
)

type thread struct {
	*base
}

func NewThread(vm *tvm.TVM) Command {
	return &thread{
		&base{
			tvm: vm,
		},
	}
}

func (c *thread) args(args ...interface{}) int {
	parentPC := toAlwaysInt(args[0])
	return parentPC
}

func (c *thread) getThreadEndAdress(memoryPos byte) int { return toAlwaysInt(memoryPos) }

func (c *thread) getParentMemory(parentThreadID int, parentPC int) []byte {
	threadManager := c.GetVM().GetThreadManager()
	parentThread := threadManager.GetThread(parentThreadID)

	parentMemory := parentThread.GetMemory()

	return parentMemory

}

func (c *thread) Execute(instruction byte, threadID int, args ...interface{}) {
	if len(args) < 2 {
		//handle error
	}

	parentPC := c.args(args...)

	parentMemory := c.getParentMemory(threadID, parentPC)
	memoryArg := parentMemory[parentPC+1]

	threadEndAdress := c.getThreadEndAdress(memoryArg)

	interpreter := c.GetVM().GetInterpreter()
	newThread := c.GetVM().GetThreadManager().NewThread(interpreter)
	threadProgram := parentMemory[parentPC+2 : parentPC+threadEndAdress]

	go newThread.Execute(threadProgram)
	c.GetCurrentThread(threadID).MovePC(threadEndAdress + 1)
}
