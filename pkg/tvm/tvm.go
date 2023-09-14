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

type Configuration struct {
	*control.ControlPlaneConfiguration
}
type TVM struct {
	cp control.ControlPlane
}

func NewTVM(c *Configuration) *TVM {
	tvm := &TVM{
		cp: control.NewControlPlane(c.ControlPlaneConfiguration),
	}

	return tvm
}

func (vm *TVM) ExecuteCode(code []byte) {
	go vm.cp.Requester()
	vm.cp.Context(0, -1, code)
}
