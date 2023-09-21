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

	current := threadManager().GetThread(threadID)

	current.SetWait()

	if current.ParentID() == -1 {
		return
	}

	c.setParentWait(current.ParentID())

}

func (c *waitThread) Execute(instruction byte, threadID int, args ...interface{}) {
	c.setParentWait(threadID)
	moveProgramPointer(1, threadID)

}
