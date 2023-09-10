package tvm

type Thread struct {
	id          int
	pc          int
	interpreter Interpreter
	operands    *operands
	memory      *memory
	heap        *heap
	stack       *stack
	pointerMap  *pointerMap
	variables   Variables
	done        chan bool
	wait        chan bool
}

func NewThread(id int, interpreter Interpreter) *Thread {
	return &Thread{
		id:          id,
		pc:          0,
		memory:      newMemory(),
		operands:    newOperands(),
		heap:        newHeap(),
		stack:       newStack(),
		pointerMap:  newPointerMap(),
		variables:   NewVariables(),
		interpreter: interpreter,
		done:        make(chan bool),
	}
}

func (t *Thread) GetVariable(name string) *Variable         { return t.variables.Get(name) }
func (t *Thread) LenVariables() int                         { return len(t.variables) }
func (t *Thread) PcPointer(pos int) int                     { return t.pc + pos }
func (t *Thread) GetObjectFromHeap(index int64) interface{} { return t.heap.Get(index) }
func (t *Thread) PopObjectFromStack() interface{}           { return t.stack.Pop() }
func (t *Thread) GetMemoryPos(pos int) byte                 { return t.memory.Get(pos) }
func (t *Thread) GetMemory() []byte                         { return t.memory.Memory() }
func (t *Thread) GetRegister(registerID byte) *register     { return t.operands.GetRegister(registerID) }
func (t *Thread) OverrideMemory(memory []byte)              { t.memory.Override(memory) }
func (t *Thread) MovePC(increment int)                      { t.pc += increment }

func (t *Thread) SetMemory(adress int, value byte)            { t.memory.Add(adress, value) }
func (t *Thread) CreateVariable(params *VariableParams)       { t.variables.NewVariable(params) }
func (t *Thread) AddToHeap(object interface{})                { t.heap.Add(object) }
func (t *Thread) PushObjectToStack(object interface{})        { t.stack.Push(object) }
func (t *Thread) RegisterInterpreter(interpreter Interpreter) { t.interpreter = interpreter }
func (t *Thread) Done() chan bool                             { return t.done }
func (t *Thread) SetDone()                                    { t.done <- true }

func (t *Thread) GetID() int { return t.id }
func (t *Thread) GetPC() int { return t.pc }

func (t *Thread) Execute(code []byte) {
	t.memory.Override(code)

	for t.pc < len(t.memory.value) {
		instruction := t.memory.value[t.pc]
		t.interpreter.Handle(instruction, t.id, t.pc)
	}
}
