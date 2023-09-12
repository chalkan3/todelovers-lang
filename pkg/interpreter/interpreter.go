package interpreter

import (
	"mary_guica/pkg/tvm"
)

const (
	ADD           = 0x00
	PRINT         = 0x01
	MOV           = 0x02
	STORE         = 0x03
	HALT          = 0x04
	LOAD_STRING   = 0x05
	CREATE_VAR    = 0x06
	LOAD_FROM_VAR = 0x07
	LOAD          = 0x08
	S_THREAD      = 0x09
	ST_THREAD     = 0x0A
	W_THREAD      = 0x0B
	SY_THREAD     = 0x0C
)

type Interpreter interface {
	Handle(instruction byte, threadID int, args ...interface{})
}
type interpreter struct {
	commands []Command
}

func NewInterpreter(vm *tvm.TVM) Interpreter {
	return &interpreter{
		commands: []Command{
			NewADD(vm),
			NewPrint(vm),
			NewMov(vm),
			NewStore(vm),
			NewHalt(vm),
			NewLoadString(vm),
			NewCreateVariable(vm),
			NewGetFromVariable(vm),
			NewLoad(vm),
			NewThread(vm),
			NewStopThread(vm),
			NewWaitThread(vm),
		},
	}
}

func (h *interpreter) Handle(instruction byte, threadID int, args ...interface{}) {
	h.commands[instruction].Execute(instruction, threadID, args...)
}
