package tvm

import (
	"fmt"
	"mary_guica/pkg/tvm/internal/memory"
	"unsafe"
)

type types byte

const (
	INT types = iota
	STRING
)

func (t types) Byte() byte {
	return [...]byte{0x00, 0x01}[t]
}

type TVM struct {
	interpreter   Interpreter
	threadManager *ThreadManager
	memoryManager memory.MemoryManager
}

func NewTVM() *TVM {
	tvm := &TVM{
		threadManager: NewThreadManager(),
		memoryManager: memory.NewMemoryManager(memory.NewMemoryAllocator(1024)),
	}

	return tvm
}

func (vm *TVM) RegisterInterpreter(interpreter Interpreter) { vm.interpreter = interpreter }
func (vm *TVM) GetThreadManager() *ThreadManager            { return vm.threadManager }
func (vm *TVM) GetInterpreter() Interpreter                 { return vm.interpreter }

func (vm *TVM) LoadCode(code []byte) {
	vm.ExecuteCode(code)
}

func (vm *TVM) ExecuteCode(code []byte) {
	mainThread := vm.threadManager.NewThread(vm.interpreter, -1)
	mainThread.Next()

	go vm.GetThreadManager().Manage()
	mainThread.Execute(code)

}

func Teste() {
	manager := memory.NewMemoryManager(memory.NewMemoryAllocator(1024))
	// manager.MapPage(0x00, 0x00)
	manager.AllocateHeap(1024)

	a := "que doideira e essa "
	b := "puts loucura"

	ptr := manager.Malloc(20)
	ptr2 := manager.Malloc(20)

	manager.Memcpy(ptr, unsafe.Pointer(&a), len(a))
	manager.Memcpy(ptr2, unsafe.Pointer(&b), len(b))

	fmt.Println(*(*string)(ptr))
	fmt.Println(*(*string)(ptr2))
	fmt.Println(*(*string)(ptr))
	// manager.Free(ptr)
	// manager.Free(ptr2)

	manager.MapPage(0x0A, 0x0A)
	manager.MapPage(0x0B, 0x0B)
	manager.MapPage(0x0C, 0x0C)
	id := manager.NewStack()

	// manager.Push(id, 0x01)
	// manager.Push(id, 0x03)

	v1 := manager.Pop(id)
	v2 := manager.Pop(id)

	fmt.Println(v1, v2)

}
