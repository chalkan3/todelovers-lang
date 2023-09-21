package threads

import (
	"mary_guica/pkg/nando"
	eapi "mary_guica/pkg/tvm/internal/api/events"
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

type Thread struct {
	metadata       *Metadata
	id             int
	events         events.EventController
	programPointer int
	parentID       int
	state          *TState
	action         *Controll
	eventAPIClient *nando.Client
}

func NewThread(id int, parentID int) *Thread {
	t := &Thread{
		metadata:       NewMetadata(id, parentID),
		eventAPIClient: eapi.Client(),
		id:             id,
		parentID:       parentID,
		action: &Controll{
			Done: make(chan bool, 1),
			Next: make(chan bool, 1),
			Wait: &WaitThread{
				Freeze:  make(chan bool, 1),
				Release: make(chan bool, 1),
			},
		},
	}

	t.state = NewTState(t)
	return t

}

func (t *Thread) Next()                  { t.action.Next <- true }
func (t *Thread) SetWait()               { t.action.Wait.Freeze <- true }
func (t *Thread) SetWaitRelease()        { t.action.Wait.Release <- true }
func (t *Thread) Wait() chan bool        { return t.action.Wait.Freeze }
func (t *Thread) ParentID() int          { return t.metadata.ParentID() }
func (t *Thread) WaitRelease() chan bool { return t.action.Wait.Release }
func (t *Thread) Done() chan bool        { return t.action.Done }
func (t *Thread) SetDone() {
	t.action.Done <- true
}
func (t *Thread) GetID() int { return t.id }
func (t *Thread) Execute(run func(threadID int, args ...interface{}), threadID int, args ...interface{}) {
	for {
		select {
		case <-t.Done():
			t.state.SetState(NewDoneState())
			t.state.Do()
			return
		case <-t.Wait():
			t.state.SetState(NewWaitState())
			t.state.Do()
		case <-t.WaitRelease():
			t.Next()
		case <-t.action.Next:
			t.state.SetState(NewRunningState())
			t.state.Do()
			run(threadID, args...)
		}

	}

}
