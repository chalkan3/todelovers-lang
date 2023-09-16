package runtime

type RunnerController interface {
	RunnerManager() RunnerManager
}
type runnerController struct {
	manager RunnerManager
}

func NewRunnerController() RunnerController {
	return &runnerController{
		manager: NewRunnerManager(),
	}
}

func (m *runnerController) RunnerManager() RunnerManager {
	return m.manager
}
