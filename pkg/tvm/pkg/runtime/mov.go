package runtime

type mov struct {
	*base
}

func NewMov(r FlightAttendant) Command {
	return &mov{
		&base{
			requester: r,
		},
	}
}

func (c *mov) Execute(instruction byte, threadID int, args ...interface{}) {
	// arg1 := c.GetCurrentThread(threadID).PcPointer(1)
	// arg2 := c.GetCurrentThread(threadID).PcPointer(2)

	// fromRegisterID := c.GetCurrentThread(threadID).GetMemoryPos(arg1)
	// toRegisterID := c.GetCurrentThread(threadID).GetMemoryPos(arg2)

	// fromReg := c.GetCurrentThread(threadID).GetRegister(fromRegisterID)
	// toReg := c.GetCurrentThread(threadID).GetRegister(toRegisterID)

	// toReg.Set(fromReg.Value())
	// fromReg.Set(1)

	// c.GetCurrentThread(threadID).MovePC(3)
}
