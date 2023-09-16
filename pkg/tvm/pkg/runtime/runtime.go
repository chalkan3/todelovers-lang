package runtime

import (
	"log"
	control "mary_guica/pkg/tvm/pkg/control_plane"
	"mary_guica/pkg/tvm/pkg/events"
	"mary_guica/pkg/tvm/pkg/metrics"
)

type Runtime interface {
	Context(id int, parent int, code []byte)
	Startup()
	EventController() events.EventController
	ControlPlane() control.ControlPlane
	Update()
}
type runtime struct {
	cp   control.ControlPlane
	ev   events.EventController
	mc   metrics.MetricsController
	rc   RunnerController
	crew Crew
}

func (rt *runtime) Context(id int, parent int, code []byte) {
	rt.cp.ProgramManager().NewFork(byte(id), code)
	rt.crew.Register(id)
	rt.ev.Notify("NEW_CREW")
	flightAttendant := rt.crew.Get(id)
	runner := rt.rc.RunnerManager().NewRunner(id, flightAttendant)
	thread := rt.cp.ThreadManager().NewThread(id, parent)
	thread.Next()
	thread.Execute(runner.Run)

}

func (rt *runtime) EventController() events.EventController { return rt.ev }
func (rt *runtime) ControlPlane() control.ControlPlane      { return rt.cp }

func (rt *runtime) Update() {
	log.Println("[VM] Updating runtime")
}

func (rt *runtime) Startup() {
	go rt.mc.GorotinesManger().Count()
	go rt.ev.Listen()
	rt.registerEvents()

}

func (rt *runtime) Requester() {

}

func (rt *runtime) registerEvents() {
	rt.ev.NewHandler("NEW_CREW")
	rt.ev.NewEvent("NEW_CREW", rt.crew)

}

func NewRuntime(cp control.ControlPlane) Runtime {
	run := &runtime{
		cp: cp,
		ev: events.NewEventController(),
		mc: metrics.NewMetricsController(),
		rc: NewRunnerController(),
	}
	run.crew = NewCrew(run)

	return run

}
