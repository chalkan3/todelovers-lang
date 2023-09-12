package commands

import "mary_guica/pkg/tvm"

type Command interface {
	Execute(instruction byte, threadID int, args ...interface{})
}

type base struct {
	tvm *tvm.TVM
}

func (b *base) SetVM(tvm *tvm.TVM) { b.tvm = tvm }
func (b *base) GetVM() *tvm.TVM    { return b.tvm }
func (b *base) GetCurrentThread(id int) *tvm.Thread {
	return b.GetVM().GetThreadManager().GetThread(id)
}

func (b *base) GetThreadID(id int) *tvm.Thread {
	return b.GetVM().GetThreadManager().GetThread(id)
}
