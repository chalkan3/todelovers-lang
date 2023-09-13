package commands

import "mary_guica/pkg/tvm"

type loadString struct {
	*base
}

func NewLoadString(vm *tvm.TVM) Command {
	return &loadString{
		&base{
			tvm: vm,
		},
	}
}

func (c *loadString) Execute(instruction byte, threadID int, args ...interface{}) {
	// arg1 := c.GetCurrentThread(threadID).PcPointer(1)
	// arg2 := c.GetCurrentThread(threadID).PcPointer(2)

	// registerID := c.GetCurrentThread(threadID).GetMemoryPos(arg1)
	// register := c.GetCurrentThread(threadID).GetRegister(registerID)

	// length := c.GetCurrentThread(threadID).GetMemoryPos(arg2)

	// strData := make([]byte, length)

	// for i := byte(0); i < length; i++ {
	// 	n := byte(c.GetCurrentThread(threadID).PcPointer(3)) + i
	// 	strData[i] = c.GetCurrentThread(threadID).GetMemoryPos(int(n))
	// }

	// str := string(strData)

	// register.Set(str)

	// c.GetCurrentThread(threadID).MovePC(3 + int(length))
}
