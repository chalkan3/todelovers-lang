package commands

import "mary_guica/pkg/tvm"

type load struct {
	*base
}

func NewLoad(vm *tvm.TVM) Command {
	return &load{
		&base{
			tvm: vm,
		},
	}
}

func (c *load) Execute(instruction byte, threadID int, args ...interface{}) {
	// arg1 := c.GetCurrentThread(threadID).PcPointer(1)
	// arg2 := c.GetCurrentThread(threadID).PcPointer(2)

	// address := c.GetCurrentThread(threadID).GetMemoryPos(arg1)
	// registerID := c.GetCurrentThread(threadID).GetMemoryPos(arg2)

	// register := c.GetCurrentThread(threadID).GetRegister(registerID)
	// register.Set(address)

	// c.GetCurrentThread(threadID).MovePC(3)
}
