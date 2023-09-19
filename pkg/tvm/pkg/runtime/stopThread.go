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

	current := c.Request(func(rt Runtime) interface{} {
		return rt.ControlPlane().ThreadManager().GetThread(threadID)
	}).ToThread()

	if current.ParentID() == -1 {
		return
	}

	parent := c.Request(func(rt Runtime) interface{} {
		return rt.ControlPlane().ThreadManager().GetParent(threadID)
	}).ToThread()

	parent.SetWaitRelease()

	// if parent.Waiting() {
	// 	parent.SetWaitRelease()
	// 	c.setParentDont(parent.GetID())
	// }

}

func (c *stopThread) Execute(instruction byte, threadID int, args ...interface{}) {
	c.setParentDont(threadID)
	c.Request(func(pm Runtime) interface{} {
		pm.ControlPlane().ThreadManager().GetThread(threadID).SetDone()
		return nil
	})

}
