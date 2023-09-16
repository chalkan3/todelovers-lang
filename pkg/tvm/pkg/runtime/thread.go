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

	// parentPC := c.args(args...)
	// parentThreadID := c.Request(func(pm control.ControlPlane) interface{} {
	// 	return pm.ThreadManager().GetParent(threadID)
	// }).ToInt()

	// fmt.Println(parentThreadID)
	// parentMemory := c.getParentMemory(threadID, parentPC)
	// memoryArg := parentMemory[parentPC+1]

	// threadEndAdress := c.getThreadEndAdress(memoryArg)

	// interpreter := c.GetVM().GetInterpreter()
	// newThread := c.GetVM().GetThreadManager().NewThread(interpreter, threadID)
	// threadProgram := parentMemory[parentPC+2 : parentPC+threadEndAdress+3]

	// newThread.Next()
	// c.GetCurrentThread(threadID).MovePC(threadEndAdress + 3)
	// go newThread.Execute(threadProgram)

}
