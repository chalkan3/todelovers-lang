package interpreter

import (
	"mary_guica/pkg/interpreter/internal/commands"
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
			commands.NewADD(vm),
			commands.NewPrint(vm),
			commands.NewMov(vm),
			commands.NewStore(vm),
			commands.NewHalt(vm),
			commands.NewLoadString(vm),
			commands.NewCreateVariable(vm),
			commands.NewGetFromVariable(vm),
			commands.NewLoad(vm),
			commands.NewThread(vm),
			commands.NewStopThread(vm),
			commands.NewWaitThread(vm),
		},
	}
}

func (h *interpreter) Handle(instruction byte, threadID int, args ...interface{}) {
	h.commands[instruction].Execute(instruction, threadID, args...)
}
