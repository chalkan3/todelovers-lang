package runtime

type stopThread struct {
	*base
}

func NewStopThread(r FlightAttendant) Command {
	return &stopThread{
		&base{
			requester: r,
		},
	}
}

func (c *stopThread) setParentDont(threadID int) {

	current := threadManager().GetThread(threadID)

	if current.ParentID() == -1 {
		return
	}

	parent := threadManager().GetParent(threadID)
	parent.SetWaitRelease()

	// if parent.Waiting() {
	// 	parent.SetWaitRelease()
	// 	c.setParentDont(parent.GetID())
	// }

}

func (c *stopThread) Execute(instruction byte, threadID int, args ...interface{}) {
	c.setParentDont(threadID)
	threadManager().GetThread(threadID).SetDone()
}
