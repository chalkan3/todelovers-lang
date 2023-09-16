package runtime

type load struct {
	*base
}

func NewLoad(r FlightAttendant) Command {
	return &load{
		&base{
			requester: r,
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
