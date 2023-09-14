package runner

import (
	"mary_guica/pkg/interpreter"
	control "mary_guica/pkg/tvm/pkg/control_plane/requester"
)

type RunnerManager interface {
	NewRunner(id int, f control.FlightAttendant, i interpreter.Interpreter) Runner
}
type runnerManager struct {
	runners []Runner
}

func NewRunnerManager() RunnerManager {
	return &runnerManager{
		runners: []Runner{},
	}
}

func (m *runnerManager) NewRunner(id int, f control.FlightAttendant, i interpreter.Interpreter) Runner {
	runner := NewRunner(id, f, i)
	m.runners = append(m.runners, runner)
	return runner
}
