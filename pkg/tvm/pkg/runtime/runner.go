package runtime

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
	flightAttendant FlightAttendant
	interpreter     Interpreter
}

func NewRunner(id int, f FlightAttendant) Runner {
	return &runner{
		id: &metadata{
			id:      id,
			running: true,
			status:  true,
		},
		flightAttendant: f,
		interpreter:     NewInterpreter(f),
	}
}

func (m *runner) Run(threadID int, args ...interface{}) {

	m.flightAttendant.Request(func(pm Runtime) interface{} {
		return pm.ControlPlane().ProgramManager().Instruction(byte(threadID))
	})
	instruction := <-m.flightAttendant.WaitForReponse()

	m.interpreter.Handle(instruction.ToByte(), threadID, args...)
}
