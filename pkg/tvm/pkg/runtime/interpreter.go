package runtime

const (
	ADD byte = iota
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

func NewInterpreter(r FlightAttendant) Interpreter {
	return &interpreter{
		commands: []Command{
			NewADD(r),
			NewPrint(r),
			NewMov(r),
			NewStore(r),
			NewHalt(r),
			NewLoadString(r),
			NewCreateVariable(r),
			NewGetFromVariable(r),
			NewLoad(r),
			NewThread(r),
			NewStopThread(r),
			NewWaitThread(r),
		},
	}
}

func (h *interpreter) Handle(instruction byte, threadID int, args ...interface{}) {
	h.commands[instruction].Execute(instruction, threadID, args...)
}
