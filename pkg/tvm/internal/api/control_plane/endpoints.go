package cplaneapi

import (
	"mary_guica/pkg/nando"
	"mary_guica/pkg/tvm/pkg/memory"
	"mary_guica/pkg/tvm/pkg/program"
	"mary_guica/pkg/tvm/pkg/register"
	"mary_guica/pkg/tvm/pkg/threads"
)

type Endpoint func(req *nando.Request) (*nando.Response, error)
type ControlPlaneAPIResponse interface {
	GetManager() interface{}
}

type ThreadManagerRequest struct{}
type ThreadManagerResponse struct {
	Manager threads.ThreadManager
}

func threadManagerRequest(req *nando.Request) *ThreadManagerRequest {
	return req.Data.(*ThreadManagerRequest)
}

func ThreadManagerEndpoint(svc Service) Endpoint {
	return func(req *nando.Request) (*nando.Response, error) {
		// _ = threadManagerRequest(req)
		threadManager, err := svc.GetThreadManager()
		return &nando.Response{
			Data: &ThreadManagerResponse{
				Manager: threadManager,
			},
		}, err
	}
}

type RegisterManagerRequest struct{}
type RegisterManagerResponse struct {
	Manager register.RegisterManager
}

func registerManagerRequest(req *nando.Request) *RegisterManagerRequest {
	return req.Data.(*RegisterManagerRequest)
}

func RegisterManagerEndpoint(svc Service) Endpoint {
	return func(req *nando.Request) (*nando.Response, error) {
		// _ = memoryManagerRequest(req)
		registerManager, err := svc.GetRegisterManager()
		return &nando.Response{
			Data: &RegisterManagerResponse{
				Manager: registerManager,
			},
		}, err
	}
}

type MemoryManagerRequest struct{}
type MemoryManagerResponse struct {
	Manager memory.MemoryManager
}

func memoryManagerRequest(req *nando.Request) *MemoryManagerRequest {
	return req.Data.(*MemoryManagerRequest)
}

func MemoryManagerEndpoint(svc Service) Endpoint {
	return func(req *nando.Request) (*nando.Response, error) {
		// _ = memoryManagerRequest(req)
		memoryManager, err := svc.GetMemoryManager()
		return &nando.Response{
			Data: &MemoryManagerResponse{
				Manager: memoryManager,
			},
		}, err
	}
}

type ProgramManagerRequest struct{}
type ProgramManagerResponse struct {
	Manager program.ProgramManager
}

func programManagerRequest(req *nando.Request) *ProgramManagerRequest {
	return req.Data.(*ProgramManagerRequest)
}

func ProgramManagerEndpoint(svc Service) Endpoint {
	return func(req *nando.Request) (*nando.Response, error) {
		// _ = programManagerRequest(req)
		programManager, err := svc.GetProgramManager()
		return &nando.Response{
			Data: &ProgramManagerResponse{
				Manager: programManager,
			},
		}, err
	}
}

func (r *ThreadManagerResponse) GetManager() interface{} {
	return r.Manager
}
func (r *MemoryManagerResponse) GetManager() interface{} {
	return r.Manager
}
func (r *RegisterManagerResponse) GetManager() interface{} {
	return r.Manager
}
func (r *ProgramManagerResponse) GetManager() interface{} {
	return r.Manager
}

func Cast[T any](t *nando.Response) T {
	return t.Data.(ControlPlaneAPIResponse).GetManager().(T)
}
