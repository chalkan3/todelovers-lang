package tvm

import (
	eapi "mary_guica/pkg/tvm/internal/api/events"
	control "mary_guica/pkg/tvm/pkg/control_plane"

	rt "mary_guica/pkg/tvm/pkg/runtime"
)

type types byte

const (
	INT types = iota
	STRING
)

func (t types) Byte() byte {
	return [...]byte{0x00, 0x01}[t]
}

type MemoryManagerConfig = control.MemoryManagerConfig
type ThreadManagerConfig = control.ThreadManagerConfig
type ProgramManagerConfig = control.ProgramManagerConfig
type ControlPlaneConfiguration = control.ControlPlaneConfiguration

type TVM struct {
	runtime  rt.Runtime
	eventAPI eapi.EventsAPI
}

func NewTVM(c *ControlPlaneConfiguration) *TVM {
	tvm := &TVM{
		runtime:  rt.NewRuntime(control.NewControlPlane(c)),
		eventAPI: eapi.NewEventsAPI(),
	}

	return tvm
}

func (vm *TVM) ExecuteCode(code []byte) {
	go vm.eventAPI.Serve()
	vm.runtime.Startup()
	vm.runtime.Context(0, -1, code)
}
