package tvm

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
	id         int
	pc         int
	parentID   int
	heap       *heap
	pointerMap *pointerMap
	variables  Variables
	state      ThreadState
	action     *Controll
}

func NewThread(id int, parentID int) *Thread {

	return &Thread{
		id:         id,
		pc:         0,
		parentID:   parentID,
		heap:       newHeap(),
		pointerMap: newPointerMap(),
		variables:  NewVariables(),
		state:      STHREAD_IDDLE,
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

func (t *Thread) CreateVariable(params *VariableParams) { t.variables.NewVariable(params) }
func (t *Thread) AddToHeap(object interface{})          { t.heap.Add(object) }
func (t *Thread) Done() chan bool                       { return t.action.Done }
func (t *Thread) SetDone() {
	t.action.Done <- true
}

func (t *Thread) GetID() int { return t.id }
func (t *Thread) GetPC() int { return t.pc }

func (t *Thread) Execute(code []byte) {

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

		}

	}

}
