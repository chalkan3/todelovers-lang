package runtime

import "mary_guica/pkg/cast"

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
	registerID := getArgument(1, threadID)
	stringLen := cast.ToAlwaysInt(getArgument(2, threadID))

	register := getRegisterByID(registerID)

	strData := make([]byte, stringLen)

	for i := 0; i < stringLen; i++ {
		strData[i] = getArgument(3+i, threadID)
	}

	str := string(strData)
	register.Set(str)

	moveProgramPointer(stringLen+3, threadID)
}
