package runtime

import (
	"fmt"
	"mary_guica/pkg/nando"
	capi "mary_guica/pkg/tvm/internal/api/control_plane"
	control "mary_guica/pkg/tvm/pkg/control_plane"
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

func getThreadManager() threads.ThreadManager {
	return getManager[threads.ThreadManager]("thread.manager")
}

func (b *base) Request(fn interface{}) Output {
	threadManager := getThreadManager()
	fmt.Println(threadManager)
	go b.requester.Request(fn)
	return <-b.requester.WaitForReponse()
}

func (b *base) MoveProgramPointer(m int, threadID int) {
	b.Request(func(ct Runtime) interface{} {
		if m == 1 {
			ct.ControlPlane().ProgramManager().Next(byte(threadID))
			return nil
		}

		ct.ControlPlane().ProgramManager().Jump(m, byte(threadID))

		ct.ControlPlane().ThreadManager().GetThread(threadID).Next()

		return nil
	})

}

func (b *base) GetArgument(pos int, threadID int) Output {
	return b.Request(func(ct Runtime) interface{} {
		return ct.ControlPlane().ProgramManager().GetAdressValue(pos, byte(threadID))
	})

}

func (b *base) GetRegisterByID(id byte) register.Register {
	return b.Request(func(ct Runtime) interface{} {
		return ct.ControlPlane().RegisterManager().GetRegister(id)
	}).ToRegister()
}

func (b *base) GetCurrentPC(threadID int) int {
	return b.Request(func(ct Runtime) interface{} {
		return ct.ControlPlane().ProgramManager().Current(byte(threadID))
	}).ToInt()
}
