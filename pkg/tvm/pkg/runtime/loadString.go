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

	c.MoveProgramPointer(stringLen+3, threadID)
}
