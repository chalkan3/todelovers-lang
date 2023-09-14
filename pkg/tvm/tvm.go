package tvm

import (
	control "mary_guica/pkg/tvm/pkg/control_plane"
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
	cp control.ControlPlane
}

func NewTVM(c *ControlPlaneConfiguration) *TVM {
	tvm := &TVM{
		cp: control.NewControlPlane(c),
	}

	return tvm
}

func (vm *TVM) ExecuteCode(code []byte) {
	go vm.cp.Requester()
	vm.cp.Context(1, 0, code)
}
