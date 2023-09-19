package runtime

type thread struct {
	*base
}

func NewThread(r FlightAttendant) Command {
	return &thread{
		&base{
			requester: r,
		},
	}
}

func (c *thread) args(args ...interface{}) int {
	parentPC := toAlwaysInt(args[0])
	return parentPC
}

func (c *thread) getThreadEndAdress(memoryPos byte) int { return toAlwaysInt(memoryPos) }

// func (c *thread) getParentMemory(parentThreadID int, parentPC int) []byte {
// 	threadManager := c.GetVM().GetThreadManager()
// 	parentThread := threadManager.GetThread(parentThreadID)

// 	parentMemory := parentThread.GetMemory()

// 	return parentMemory

// }

func (c *thread) Execute(instruction byte, threadID int, args ...interface{}) {

	parentThreadID := threadID
	endThread := c.GetArgument(1, threadID).ToInt()
	currentPC := c.GetCurrentPC(threadID)

	threadCode := c.Request(func(rt Runtime) interface{} {
		code := rt.ControlPlane().ProgramManager().Code(byte(threadID))
		c := code[currentPC+2 : currentPC+3+endThread]
		return c
	}).ToByteArray()

	c.Request(func(rt Runtime) interface{} {
		go rt.Context(parentThreadID+1, parentThreadID, threadCode)
		return nil
	})

	c.MoveProgramPointer(3+endThread, threadID)

}
