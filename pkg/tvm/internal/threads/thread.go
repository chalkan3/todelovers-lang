package threads

type WaitThread struct {
	Freeze  chan bool
	Release chan bool
}
type Controll struct {
	Next chan bool
	Done chan bool
	Wait *WaitThread
}

type ThreadState int

func (ts ThreadState) String() string {
	return [...]string{"WAIT", "DONE", "RUNNING", "IDDLE"}[ts]
}

const (
	STHREAD_WAIT ThreadState = iota
	STHREAD_DONE
	STHREAD_RUNNING
	STHREAD_IDDLE
)

type Thread struct {
	id          int
	pc          int
	parentID    int
	interpreter Interpreter
	operands    *operands
	memory      *memory
	state       ThreadState
	action      *Controll
}

func NewThread(id int, parentID int, interpreter Interpreter) *Thread {

	return &Thread{
		id:          id,
		pc:          0,
		parentID:    parentID,
		memory:      NewMemory(),
		operands:    newOperands(),
		interpreter: interpreter,
		state:       STHREAD_IDDLE,
		action: &Controll{
			Done: make(chan bool, 1),
			Next: make(chan bool, 1),
			Wait: &WaitThread{
				Freeze:  make(chan bool, 1),
				Release: make(chan bool, 1),
			},
		},
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
func (t *Thread) Next()                                     { t.action.Next <- true }
func (t *Thread) Wait() chan bool                           { return t.action.Wait.Freeze }
func (t *Thread) Waiting() bool                             { return t.state == STHREAD_WAIT }

func (t *Thread) ParentID() int { return t.parentID }
func (t *Thread) SetWait() {
	t.action.Wait.Freeze <- true
}
func (t *Thread) WaitRelease() chan bool { return t.action.Wait.Release }

func (t *Thread) SetWaitRelease() {
	t.action.Wait.Release <- true
}

func (t *Thread) MovePC(increment int) {
	defer t.Next()
	t.pc += increment

}

func (t *Thread) SetMemory(adress int, value byte)            { t.memory.Add(adress, value) }
func (t *Thread) CreateVariable(params *VariableParams)       { t.variables.NewVariable(params) }
func (t *Thread) AddToHeap(object interface{})                { t.heap.Add(object) }
func (t *Thread) PushObjectToStack(object interface{})        { t.stack.Push(object) }
func (t *Thread) RegisterInterpreter(interpreter Interpreter) { t.interpreter = interpreter }
func (t *Thread) Done() chan bool                             { return t.action.Done }
func (t *Thread) SetDone() {
	t.action.Done <- true
}

func (t *Thread) GetID() int { return t.id }
func (t *Thread) GetPC() int { return t.pc }

func (t *Thread) Execute(code []byte) {
	t.memory.Override(code)

	for {
		select {
		case <-t.Done():
			t.state = STHREAD_DONE
			return
		case <-t.Wait():
			t.state = STHREAD_WAIT
		case <-t.WaitRelease():
			t.MovePC(1)
		case <-t.action.Next:
			t.state = STHREAD_RUNNING
			instruction := t.memory.value[t.pc]
			t.interpreter.Handle(instruction, t.id, t.pc)
		}

	}

}
