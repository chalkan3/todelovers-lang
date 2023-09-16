package runtime

type loadString struct {
	*base
}

func NewLoadString(r FlightAttendant) Command {
	return &loadString{
		&base{
			requester: r,
		},
	}
}

func (c *loadString) Execute(instruction byte, threadID int, args ...interface{}) {
	registerID := c.GetArgument(1, threadID).ToByte()
	stringLen := c.GetArgument(2, threadID).ToInt()

	register := c.GetRegisterByID(registerID)

	strData := make([]byte, stringLen)

	for i := 0; i < stringLen; i++ {
		strData[i] = c.GetArgument(3+i, threadID).ToByte()
	}

	str := string(strData)
	register.Set(str)

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
	c.MoveProgramPointer(stringLen+3, threadID)
}
