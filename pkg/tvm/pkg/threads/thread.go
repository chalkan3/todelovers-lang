package threads

import (
	"fmt"
	"log"
	"mary_guica/pkg/tvm/pkg/events"
)

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
	metadata       *Metadata
	id             int
	events         events.EventController
	programPointer int
	parentID       int
	state          ThreadState
	action         *Controll
}

func NewThread(id int, parentID int, ec events.EventController) *Thread {
	t := &Thread{
		metadata: NewMetadata(id, parentID),
		id:       id,
		parentID: parentID,
		events:   ec,
		state:    STHREAD_IDDLE,
		action: &Controll{
			Done: make(chan bool, 1),
			Next: make(chan bool, 1),
			Wait: &WaitThread{
				Freeze:  make(chan bool, 1),
				Release: make(chan bool, 1),
			},
		},
	}

	ec.NewHandler(fmt.Sprintf("EVENT_THREAD_ID_%d", id))

	ec.NewEvent(fmt.Sprintf("EVENT_THREAD_ID_%d", id), t)
	return t

}

func (t *Thread) Next()                      { t.action.Next <- true }
func (t *Thread) SetWait()                   { t.action.Wait.Freeze <- true }
func (t *Thread) SetWaitRelease()            { t.action.Wait.Release <- true }
func (t *Thread) MoveProgramPointer(pos int) { t.programPointer += pos }
func (t *Thread) GetProgramArg(pos int) int  { return t.programPointer + pos }
func (t *Thread) Wait() chan bool            { return t.action.Wait.Freeze }
func (t *Thread) Waiting() bool              { return t.state == STHREAD_WAIT }
func (t *Thread) ParentID() int              { return t.metadata.ParentID() }
func (t *Thread) WaitRelease() chan bool     { return t.action.Wait.Release }

func (t *Thread) MovePC(increment int) {
	defer t.Next()
	// t.pc += increment
}

func (t *Thread) Done() chan bool { return t.action.Done }
func (t *Thread) SetDone() {
	t.action.Done <- true
}

func (t *Thread) GetID() int { return t.id }

func (t *Thread) GetPC() int { return 1 } //t.pc }

func (t *Thread) Update() {
	log.Printf("[VM] Thread ID: %d is ended\n", t.metadata.id)

}

func (t *Thread) Execute(run func(threadID int, args ...interface{}), threadID int, args ...interface{}) {
	for {
		select {
		case <-t.Done():
			t.state = STHREAD_DONE
			t.events.Notify(fmt.Sprintf("EVENT_THREAD_ID_%d", t.metadata.id))
			return
		case <-t.Wait():
			t.state = STHREAD_WAIT
		case <-t.WaitRelease():
			t.MovePC(1)
		case <-t.action.Next:
			t.state = STHREAD_RUNNING
			run(threadID, args...)
		}

	}

}
