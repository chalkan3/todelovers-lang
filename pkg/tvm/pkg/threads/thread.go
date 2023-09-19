package threads

import (
	"fmt"
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

func NewThread(id int, parentID int) *Thread {
	t := &Thread{
		metadata: NewMetadata(id, parentID),
		id:       id,
		parentID: parentID,
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

func (t *Thread) Execute(run func(threadID int, args ...interface{}), threadID int, args ...interface{}) {
	for {
		select {
		case <-t.Done():
			t.state = STHREAD_DONE
			events.GetEventController().Notify(&events.Notifier{
				Handler: "NOTIFY",
				Event: &events.Event{
					Name:        "THREAD_DONE",
					Description: fmt.Sprintf("thread_ID :[%d]", t.metadata.id),
					Data:        nil,
				},
			})
			return
		case <-t.Wait():
			t.state = STHREAD_WAIT
			events.GetEventController().Notify(&events.Notifier{
				Handler: "NOTIFY",
				Event: &events.Event{
					Name:        "THREAD_WAIT",
					Description: fmt.Sprintf("thread_ID :[%d]", t.metadata.id),
					Data:        nil,
				},
			})
		case <-t.WaitRelease():
			events.GetEventController().Notify(&events.Notifier{
				Handler: "NOTIFY",
				Event: &events.Event{
					Name:        "RELEASED",
					Description: fmt.Sprintf("thread_ID :[%d]", t.metadata.id),
					Data:        nil,
				},
			})
			t.Next()
		case <-t.action.Next:
			t.state = STHREAD_RUNNING
			run(threadID, args...)
		}

	}

}
