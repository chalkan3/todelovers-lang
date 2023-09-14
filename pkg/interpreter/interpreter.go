package interpreter

import (
	"mary_guica/pkg/interpreter/internal/commands"
	control "mary_guica/pkg/tvm/pkg/control_plane/requester"
)

const (
	ADD byte = iota + 1
	PRINT
	MOV
	STORE
	HALT
	LOAD_STRING
	CREATE_VAR
	LOAD_FROM_VAR
	LOAD
	S_THREAD
	ST_THREAD
	W_THREAD
	SY_THREAD
	CALL
	RET
)

type Interpreter interface {
	Handle(instruction byte, threadID int, args ...interface{})
}

type interpreter struct {
	commands []Command
}

func NewInterpreter(r control.FlightAttendant) Interpreter {
	return &interpreter{
		commands: []Command{
			commands.NewADD(r),
			commands.NewPrint(r),
			commands.NewMov(r),
			commands.NewStore(r),
			commands.NewHalt(r),
			commands.NewLoadString(r),
			commands.NewCreateVariable(r),
			commands.NewGetFromVariable(r),
			commands.NewLoad(r),
			commands.NewThread(r),
			commands.NewStopThread(r),
			commands.NewWaitThread(r),
		},
	}
}

func (h *interpreter) Handle(instruction byte, threadID int, args ...interface{}) {
	h.commands[instruction].Execute(instruction, threadID, args...)
}
