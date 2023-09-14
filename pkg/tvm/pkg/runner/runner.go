package runner

import (
	"mary_guica/pkg/interpreter"
	requester "mary_guica/pkg/tvm/pkg/control_plane/requester"
	"mary_guica/pkg/tvm/pkg/program"
)

type Runner interface {
	Run(threadID int, args ...interface{})
}

type metadata struct {
	id      int
	running bool
	status  bool // change status
}
type runner struct {
	id              *metadata
	flightAttendant requester.FlightAttendant
	interpreter     interpreter.Interpreter
}

func NewRunner(id int, f requester.FlightAttendant, i interpreter.Interpreter) Runner {
	return &runner{
		id: &metadata{
			id:      id,
			running: true,
			status:  true,
		},
		flightAttendant: f,
		interpreter:     i,
	}
}

func (m *runner) Run(threadID int, args ...interface{}) {
	var instruction byte
	var pc int
	m.flightAttendant.Request(func(pm program.ProgramManager) interface{} {
		// change it !!!!!!!!!!!!
		if threadID == -1 {
			instruction = pm.Instruction()
			pc = pm.Current()
		} else {
			instruction = pm.Instruction(byte(threadID))
			pc = pm.Current(byte(threadID))
		}

		return nil
	})

	m.interpreter.Handle(instruction, threadID, args...)
}
