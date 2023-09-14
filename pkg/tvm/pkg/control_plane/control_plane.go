package control

import (
	"fmt"
	"mary_guica/pkg/interpreter"
	"mary_guica/pkg/tvm/pkg/control_plane/requester"
	"mary_guica/pkg/tvm/pkg/memory"
	"mary_guica/pkg/tvm/pkg/program"
	"mary_guica/pkg/tvm/pkg/runner"
	"mary_guica/pkg/tvm/pkg/threads"
	"reflect"
	"time"
)

type MemoryManagerConfig struct {
	FrameSize int
}

type ThreadManagerConfig struct {
	FrameSize int
}

type ProgramManagerConfig struct {
	Code []byte
}

type ControlPlaneConfiguration struct {
	MemoryManager  MemoryManagerConfig
	ThreadManager  ThreadManagerConfig
	ProgramManager ProgramManagerConfig
}

type ControlPlane interface {
	ProgramManager() program.ProgramManager
	MemoryManager() memory.MemoryManager
	ThreadManager() threads.ThreadManager
	RunnerManager() runner.RunnerManager
	Context(id int, parent int, code []byte)
	Requester()
}

type controlPlane struct {
	memoryManager  memory.MemoryManager
	threadManager  threads.ThreadManager
	programManager program.ProgramManager
	runnerManager  runner.RunnerManager
	crew           requester.Crew
}

func NewControlPlane(c *ControlPlaneConfiguration) ControlPlane {
	return &controlPlane{
		memoryManager:  memory.NewMemoryManager(memory.NewMemoryAllocator(c.MemoryManager.FrameSize)),
		threadManager:  threads.NewThreadManager(),
		programManager: program.NewProgramManager(c.ProgramManager.Code),
		runnerManager:  runner.NewRunnerManager(),
		crew:           requester.NewCrew(),
	}
}

func (cp *controlPlane) ProgramManager() program.ProgramManager { return cp.programManager }
func (cp *controlPlane) MemoryManager() memory.MemoryManager    { return cp.memoryManager }
func (cp *controlPlane) ThreadManager() threads.ThreadManager   { return cp.threadManager }
func (cp *controlPlane) RunnerManager() runner.RunnerManager    { return cp.runnerManager }

func (cp *controlPlane) Context(id int, parent int, code []byte) {
	cp.programManager.NewFork(byte(id), code)
	cp.crew.Register(id)
	flightAttendant := cp.crew.Get(id)
	interpreter := interpreter.NewInterpreter(flightAttendant)
	runner := cp.RunnerManager().NewRunner(id, flightAttendant, interpreter)
	thread := cp.ThreadManager().NewThread(id, parent, runner)
	thread.Execute()

}

func (cp *controlPlane) Requester() {
	crew := cp.crew.PrepareCrew()
	for {
		select {
		case channelID := <-crew:
			fn := <-cp.crew.Get(channelID).WaitForRequest()
			v := reflect.ValueOf(fn)
			v.Call([]reflect.Value{reflect.ValueOf(cp.MemoryManager())})

		case <-time.After(2 * time.Second):
			fmt.Println("Timeout: No data received in 2 seconds.")
		}
	}
}
