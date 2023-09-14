package runtime

import (
	control "mary_guica/pkg/tvm/pkg/control_plane"
	"mary_guica/pkg/tvm/pkg/memory"
	"mary_guica/pkg/tvm/pkg/threads"
	"reflect"
)

type MemoryManager = memory.MemoryManager
type ThreadManager = threads.ThreadManager

type Runtime interface {
	Requester()
}
type runtime struct {
	call chan interface{}
	cp   control.ControlPlane
}

func NewRuntime(call chan interface{}, cp control.ControlPlane) Runtime {
	return &runtime{
		call: call,
		cp:   cp,
	}
}

func CallFunction(fn interface{}, arg interface{}) interface{} {
	fnValue := reflect.ValueOf(fn)
	argValue := reflect.ValueOf(arg)

	result := fnValue.Call([]reflect.Value{argValue.Convert(fnValue.Type().In(0))})
	return result[0].Interface()
}

func (rt *runtime) Requester() {
	for {
		fn := <-rt.call
		v := reflect.ValueOf(fn)
		v.Call([]reflect.Value{reflect.ValueOf(rt.cp.MemoryManager())})

	}
}
