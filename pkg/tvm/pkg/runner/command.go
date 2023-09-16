package runtime

import (
	"mary_guica/pkg/tvm/pkg/register"
	requester "mary_guica/pkg/tvm/pkg/runtime"
)

type Command interface {
	Execute(instruction byte, threadID int, args ...interface{})
}

type base struct {
	requester requester.FlightAttendant
}

func (b *base) Request(fn interface{}) requester.Output {
	go b.requester.Request(fn)
	return <-b.requester.WaitForReponse()
}

func (b *base) MoveProgramPointer(m int, threadID int) {
	b.Request(func(ct requester.Runtime) interface{} {
		if m == 1 {
			ct.ControlPlane().ProgramManager().Next(byte(threadID))
			return nil
		}

		ct.ControlPlane().ProgramManager().Jump(m, byte(threadID))

		ct.ControlPlane().ThreadManager().GetThread(threadID).Next()

		return nil
	})

}

func (b *base) GetArgument(pos int, threadID int) requester.Output {
	return b.Request(func(ct requester.Runtime) interface{} {
		return ct.ControlPlane().ProgramManager().GetAdressValue(pos, byte(threadID))
	})

}

func (b *base) GetRegisterByID(id byte) register.Register {
	return b.Request(func(ct requester.Runtime) interface{} {
		return ct.ControlPlane().RegisterManager().GetRegister(id)
	}).ToRegister()
}

func (b *base) GetCurrentPC(threadID int) int {
	return b.Request(func(ct requester.Runtime) interface{} {
		return ct.ControlPlane().ProgramManager().Current(byte(threadID))
	}).ToInt()
}
