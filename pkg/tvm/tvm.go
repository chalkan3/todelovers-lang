package tvm

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
}

func NewTVM() *TVM {
	tvm := &TVM{
		threadManager: NewThreadManager(),
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
	mainThread := vm.threadManager.NewThread(vm.interpreter)
	go mainThread.Execute(code)
	for {
		<-mainThread.Done()
		break
	}
}
