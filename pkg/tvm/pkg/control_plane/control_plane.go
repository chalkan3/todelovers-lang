package control

import (
	"mary_guica/pkg/tvm/pkg/memory"
	"mary_guica/pkg/tvm/pkg/program"
	"mary_guica/pkg/tvm/pkg/register"
	"mary_guica/pkg/tvm/pkg/threads"
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
	RegisterManager() register.RegisterManager
}

type controlPlane struct {
	memoryManager   memory.MemoryManager
	threadManager   threads.ThreadManager
	programManager  program.ProgramManager
	registerManager register.RegisterManager
}

func NewControlPlane(c *ControlPlaneConfiguration) ControlPlane {
	return &controlPlane{
		memoryManager:   memory.NewMemoryManager(memory.NewMemoryAllocator(c.MemoryManager.FrameSize)),
		threadManager:   threads.NewThreadManager(),
		programManager:  program.NewProgramManager(c.ProgramManager.Code),
		registerManager: register.NewRegisterManager(),
	}
}

func (cp *controlPlane) ProgramManager() program.ProgramManager    { return cp.programManager }
func (cp *controlPlane) MemoryManager() memory.MemoryManager       { return cp.memoryManager }
func (cp *controlPlane) ThreadManager() threads.ThreadManager      { return cp.threadManager }
func (cp *controlPlane) RegisterManager() register.RegisterManager { return cp.registerManager }
