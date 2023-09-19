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

func (c *waitThread) setParentWait(threadID int) {

	current := c.Request(func(rt Runtime) interface{} {
		return rt.ControlPlane().ThreadManager().GetThread(threadID)
	}).ToThread()

	current.SetWait()

	if current.ParentID() == -1 {
		return
	}

	c.setParentWait(current.ParentID())

}

func (c *waitThread) Execute(instruction byte, threadID int, args ...interface{}) {
	c.setParentWait(threadID)
	c.MoveProgramPointer(1, threadID)

}
