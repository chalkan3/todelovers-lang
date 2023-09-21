package runtime

import (
	"fmt"
	"mary_guica/pkg/nando"
	eapi "mary_guica/pkg/tvm/internal/api/events"
	control "mary_guica/pkg/tvm/pkg/control_plane"
	"mary_guica/pkg/tvm/pkg/events"
	"mary_guica/pkg/tvm/pkg/program"

	"mary_guica/pkg/tvm/pkg/threads"
)

type Inputs struct {
	Code     []byte
	ParentID int
}
type ContainerBuilder interface {
	BuildRunner()
	BuildThread(parentID int)
	BuildProgramFork(code []byte)
	BuildCrew()
	GetCTX() *Container
}

type containerBuilder struct {
	container *Container
	cp        control.ControlPlane
	crew      Crew
}

func NewContainerBuilder(id int, cp control.ControlPlane, c Crew) *containerBuilder {
	return &containerBuilder{
		container: NewContainer(id),
		cp:        cp,
		crew:      c,
	}
}

func (c *containerBuilder) BuildRunner() {
	c.container.runner = NewRunner(c.container.id, c.crew.Get(c.container.id))
}
func (c *containerBuilder) BuildThread(parentID int) {
	c.container.thread = c.cp.ThreadManager().NewThread(c.container.id, parentID)
}
func (c *containerBuilder) BuildProgramFork(code []byte) {
	c.cp.ProgramManager().NewFork(byte(c.container.id), code)
}
func (c *containerBuilder) BuildCrew() {
	c.crew.Register(c.container.id)
	client := eapi.Client()
	client.Do(nando.NewRequest(eapi.Notify.String(), &eapi.NotifyRequest{
		Notifier: &events.Notifier{
			Handler: "NEW_CREW",
			Event: &events.Event{
				Name:        "RELOAD_CREW",
				Description: fmt.Sprintf("New flightAttendent ID:[%d] Joined the crew", c.container.id),
				Data:        nil,
			},
		},
	}))
}
func (c *containerBuilder) GetCTX() *Container {
	return c.container
}

type Container struct {
	id      int
	runner  Runner
	thread  *threads.Thread
	program program.Program
	crew    Crew
}

func (c *Container) ID() int                  { return c.id }
func (c *Container) Runner() Runner           { return c.runner }
func (c *Container) Program() program.Program { return c.program }
func (c *Container) Crew() Crew               { return c.crew }
func (c *Container) Thread() *threads.Thread  { return c.thread }

func NewContainer(id int) *Container {
	return &Container{
		id: id,
	}
}

type Manager struct {
	builder ContainerBuilder
}

func NewContainerManager(builder ContainerBuilder) *Manager {
	return &Manager{
		builder: builder,
	}
}

func (m *Manager) New(i *Inputs) *Container {
	m.builder.BuildProgramFork(i.Code)
	m.builder.BuildCrew()
	m.builder.BuildRunner()
	m.builder.BuildThread(i.ParentID)
	return m.builder.GetCTX()
}
