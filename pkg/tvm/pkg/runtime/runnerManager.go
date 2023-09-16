package runtime

type RunnerManager interface {
	NewRunner(id int, f FlightAttendant) Runner
}
type runnerManager struct {
	runners []Runner
}

func NewRunnerManager() RunnerManager {
	return &runnerManager{
		runners: []Runner{},
	}
}

func (m *runnerManager) NewRunner(id int, f FlightAttendant) Runner {
	runner := NewRunner(id, f)
	m.runners = append(m.runners, runner)
	return runner
}
