package tvm

import (
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
	runtime rt.Runtime
}

func NewTVM(c *ControlPlaneConfiguration) *TVM {
	tvm := &TVM{
		runtime: rt.NewRuntime(control.NewControlPlane(c)),
	}

	return tvm
}

func (vm *TVM) Startup() *TVM {
	vm.runtime.StartAPIS()
	vm.runtime.RegisterEvents()
	vm.runtime.StartProfiler()
	return vm
}

func (vm *TVM) ExecuteCode(code []byte) {
	vm.runtime.Context(0, -1, code)
}
