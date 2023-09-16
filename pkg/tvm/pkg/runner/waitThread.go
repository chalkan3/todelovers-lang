package runtime

type waitThread struct {
	*base
}

func NewWaitThread(r FlightAttendant) Command {
	return &waitThread{
		&base{
			requester: r,
		},
	}
}

// func (c *waitThread) setParentWait(threadID int) {
// 	current := c.GetCurrentThread(threadID)
// 	current.SetWait()

// 	if current.ParentID() != -1 {
// 		c.setParentWait(current.ParentID())
// 		return

// 	}

// }

func (c *waitThread) Execute(instruction byte, threadID int, args ...interface{}) {
	// c.setParentWait(threadID)
}
