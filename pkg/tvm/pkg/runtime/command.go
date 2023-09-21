package runtime

import (
	"mary_guica/pkg/nando"
	capi "mary_guica/pkg/tvm/internal/api/control_plane"
	control "mary_guica/pkg/tvm/pkg/control_plane"
	"mary_guica/pkg/tvm/pkg/memory"
	"mary_guica/pkg/tvm/pkg/program"
	"mary_guica/pkg/tvm/pkg/register"
	"mary_guica/pkg/tvm/pkg/threads"
)

type Command interface {
	Execute(instruction byte, threadID int, args ...interface{})
}

type base struct {
	requester FlightAttendant
}

func getManager[T control.Manager](apiPath string) T {
	c := capi.Client()
	req := nando.NewRequest(apiPath, nil)
	response, _ := c.Do(req)
	return capi.Cast[T](response)
}

func threadManager() threads.ThreadManager {
	return getManager[threads.ThreadManager]("thread.manager")
}
func programManager() program.ProgramManager {
	return getManager[program.ProgramManager]("program.manager")
}
func registerManager() register.RegisterManager {
	return getManager[register.RegisterManager]("register.manager")
}
func memoryManager() memory.MemoryManager {
	return getManager[memory.MemoryManager]("memory.manager")
}

func moveProgramPointer(m int, threadID int) {
	if m == 1 {
		programManager().Next(byte(threadID))
		return
	}

	programManager().Jump(m, byte(threadID))

	threadManager().GetThread(threadID).Next()
}

func (b *base) Request(fn interface{}) Output {
	go b.requester.Request(fn)
	return <-b.requester.WaitForReponse()
}

func getArgument(pos int, threadID int) byte {
	return programManager().GetAdressValue(pos, byte(threadID))
}
func getRegisterByID(id byte) register.Register {
	return registerManager().GetRegister(id)
}

func getCurrentPC(threadID int) int {
	return programManager().Current(byte(threadID))
}
