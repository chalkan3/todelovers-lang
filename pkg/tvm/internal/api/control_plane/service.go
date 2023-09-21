package cplaneapi

import (
	control "mary_guica/pkg/tvm/pkg/control_plane"
	"mary_guica/pkg/tvm/pkg/memory"
	"mary_guica/pkg/tvm/pkg/program"
	"mary_guica/pkg/tvm/pkg/register"
	"mary_guica/pkg/tvm/pkg/threads"
)

type Service interface {
	GetMemoryManager() (memory.MemoryManager, error)
	GetThreadManager() (threads.ThreadManager, error)
	GetProgramManager() (program.ProgramManager, error)
	GetRegisterManager() (register.RegisterManager, error)
}
type service struct {
	cp control.ControlPlane
}

func NewService(cp control.ControlPlane) Service {
	return &service{
		cp: cp,
	}
}

func (s *service) GetMemoryManager() (memory.MemoryManager, error) {
	return s.cp.MemoryManager(), nil
}
func (s *service) GetThreadManager() (threads.ThreadManager, error) {
	return s.cp.ThreadManager(), nil
}
func (s *service) GetProgramManager() (program.ProgramManager, error) {
	return s.cp.ProgramManager(), nil
}
func (s *service) GetRegisterManager() (register.RegisterManager, error) {
	return s.cp.RegisterManager(), nil
}
