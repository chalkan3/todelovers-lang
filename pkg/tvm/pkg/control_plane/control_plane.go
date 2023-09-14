package control

// Verificar o paradeiro do id -1

import (
	"fmt"
	"mary_guica/pkg/interpreter"
	"mary_guica/pkg/tvm/pkg/control_plane/requester"
	"mary_guica/pkg/tvm/pkg/memory"
	"mary_guica/pkg/tvm/pkg/program"
	"mary_guica/pkg/tvm/pkg/runner"
	"mary_guica/pkg/tvm/pkg/threads"
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
	thread := cp.ThreadManager().NewThread(id, parent)
	thread.Next()
	thread.Execute(runner.Run)

}

func (cp *controlPlane) Requester() {
	// crew :=
	for {
		select {
		case channelID := <-cp.crew.PrepareCrew():
			fn := <-cp.crew.Get(channelID).WaitForRequest()

			if fn, ok := fn.(func(program.ProgramManager) interface{}); ok {
				fn(cp.ProgramManager())
			}

			if fn, ok := fn.(func(memory.MemoryManager) interface{}); ok {
				fn(cp.MemoryManager())
			}

			if fn, ok := fn.(func(threads.ThreadManager) interface{}); ok {
				fn(cp.ThreadManager())
			}

			// cp.crew.Get(channelID).Response()

		case <-time.After(2 * time.Second):
			fmt.Println("Timeout: No data received in 2 seconds.")
		}
	}
}
