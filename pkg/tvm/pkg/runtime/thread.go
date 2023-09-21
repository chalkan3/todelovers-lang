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
	endThread := getArgument(1, threadID)
	currentPC := getCurrentPC(threadID)

	code := programManager().Code(byte(threadID))
	threadCode := code[currentPC+2 : currentPC+3+int(endThread)]

	c.Request(func(rt Runtime) interface{} {
		go rt.Context(parentThreadID+1, parentThreadID, threadCode)
		return nil
	})

	moveProgramPointer(3+int(endThread), threadID)

}
