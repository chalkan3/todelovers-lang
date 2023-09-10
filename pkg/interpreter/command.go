package interpreter

import "mary_guica/pkg/tvm"

type Command interface {
	Execute(instruction byte)
}

type base struct {
	tvm *tvm.TVM
}

func (b *base) SetVM(tvm *tvm.TVM) { b.tvm = tvm }
func (b *base) GetVM() *tvm.TVM    { return b.tvm }
