package commands

import control "mary_guica/pkg/tvm/pkg/control_plane/requester"

type waitThread struct {
	*base
}

func NewWaitThread(r control.FlightAttendant) Command {
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
