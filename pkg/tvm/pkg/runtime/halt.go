package runtime

type halt struct {
	*base
}

func NewHalt(r FlightAttendant) Command {
	return &halt{
		&base{
			requester: r,
		},
	}
}

func (c *halt) Execute(instruction byte, threadID int, args ...interface{}) {
	threadManager().GetThread(threadID).SetDone()
}
