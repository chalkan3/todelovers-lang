package runtime

import (
	"mary_guica/pkg/nando"
	eapi "mary_guica/pkg/tvm/internal/api/events"
	api "mary_guica/pkg/tvm/pkg/api_manager"

	control "mary_guica/pkg/tvm/pkg/control_plane"
	"mary_guica/pkg/tvm/pkg/logger"
	"os"

	"mary_guica/pkg/tvm/pkg/events"

	"mary_guica/pkg/tvm/pkg/metrics"
)

type Runtime interface {
	Context(id int, parent int, code []byte)
	StartAPIS()
	RegisterEvents()
	StartProfiler()
	ControlPlane() control.ControlPlane
}
type runtime struct {
	cp   control.ControlPlane
	mc   metrics.MetricsController
	rc   RunnerController
	apim api.APIManager
	crew Crew
}

func (rt *runtime) Context(id int, parent int, code []byte) {
	builder := NewContainerBuilder(id, rt.cp, rt.crew)
	containerManager := NewContainerManager(builder)
	container := containerManager.New(&Inputs{
		Code:     code,
		ParentID: parent,
	})

	thread := container.Thread()
	runner := container.Runner()

	thread.Next()
	thread.Execute(runner.Run, id)

}

func (rt *runtime) ControlPlane() control.ControlPlane { return rt.cp }

func (rt *runtime) RegisterEvents() {
	rt.registerEvents()
}

func (rt *runtime) StartProfiler() {
	go rt.mc.GorotinesManger().Count()
	go rt.cp.ThreadManager().Manage()
}

func (rt *runtime) registerEvents() {
	c := eapi.Client()

	c.Do(nando.NewRequest(eapi.CreateHandler.String(), &eapi.CreateHandlerRequest{
		EventHandler: &eapi.EventHandler{
			ID:          "1",
			HandlerName: "NEW_CREW",
			Handler: []events.Observer{
				logger.NewConsoleLogObserver(os.Stdout),
				logger.NewFileLogObserver(),
				NewReloadCrewObserver(rt.crew),
			},
		},
	}))

	c.Do(nando.NewRequest(eapi.CreateHandler.String(), &eapi.CreateHandlerRequest{
		EventHandler: &eapi.EventHandler{
			ID:          "2",
			HandlerName: "NOTIFY",
			Handler: []events.Observer{
				logger.NewConsoleLogObserver(os.Stdout),
				logger.NewFileLogObserver(),
			},
		},
	}))

}

func (rt *runtime) StartAPIS() {
	rt.apim.Launch()
}

func NewRuntime(cp control.ControlPlane) Runtime {
	run := &runtime{
		cp:   cp,
		mc:   metrics.NewMetricsController(),
		rc:   NewRunnerController(),
		apim: api.NewAPIManager(cp),
	}
	run.crew = NewCrew(run)

	return run

}
