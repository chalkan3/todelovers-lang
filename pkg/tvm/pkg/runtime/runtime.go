package runtime

import (
	"fmt"
	control "mary_guica/pkg/tvm/pkg/control_plane"
	"mary_guica/pkg/tvm/pkg/events"
	"mary_guica/pkg/tvm/pkg/logger"
	"os"

	"mary_guica/pkg/tvm/pkg/metrics"
)

type Runtime interface {
	Context(id int, parent int, code []byte)
	Startup()
	ControlPlane() control.ControlPlane
}
type runtime struct {
	cp   control.ControlPlane
	mc   metrics.MetricsController
	rc   RunnerController
	crew Crew
}

func (rt *runtime) Context(id int, parent int, code []byte) {
	rt.cp.ProgramManager().NewFork(byte(id), code)
	rt.crew.Register(id)
	flightAttendant := rt.crew.Get(id)
	events.GetEventController().Notify(&events.Notifier{
		Handler: "NEW_CREW",
		Event: &events.Event{
			Name:        "RELOAD_CREW",
			Description: fmt.Sprintf("New flightAttendent ID:[%d] Joined the crew", id),
			Data:        flightAttendant,
		},
	})
	runner := rt.rc.RunnerManager().NewRunner(id, flightAttendant)
	thread := rt.cp.ThreadManager().NewThread(id, parent)
	events.GetEventController().Notify(&events.Notifier{
		Handler: "NOTIFY",
		Event: &events.Event{
			Name:        "NEW_THREAD",
			Description: fmt.Sprintf("New thread id:[%d]", thread.GetID()),
			Data:        flightAttendant,
		},
	})
	thread.Next()
	thread.Execute(runner.Run, id)

}

func (rt *runtime) ControlPlane() control.ControlPlane { return rt.cp }

func (rt *runtime) Startup() {
	// go rt.mc.GorotinesManger().Count()
	go events.GetEventController().Listen()
	go rt.cp.ThreadManager().Manage()
	rt.registerEvents()

}

func (rt *runtime) registerEvents() {
	events.GetEventController().NewObserver("NEW_CREW", []events.Observer{
		logger.NewConsoleLogObserver(os.Stdout),
		logger.NewFileLogObserver(),
		NewReloadCrewObserver(rt.crew),
	})

	events.GetEventController().NewObserver("NOTIFY", []events.Observer{
		logger.NewConsoleLogObserver(os.Stdout),
		logger.NewFileLogObserver(),
	})

}

func NewRuntime(cp control.ControlPlane) Runtime {
	run := &runtime{
		cp: cp,
		mc: metrics.NewMetricsController(),
		rc: NewRunnerController(),
	}
	run.crew = NewCrew(run)

	return run

}
