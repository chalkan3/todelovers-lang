package threads

import (
	"fmt"
	"mary_guica/pkg/nando"
	eapi "mary_guica/pkg/tvm/internal/api/events"
	"mary_guica/pkg/tvm/pkg/events"
)

type State interface {
	Next() State
	Back() State
	Do(thread *Thread)
	Set(state State) State
}
type TState struct {
	thread  *Thread
	current State
	history []State
}

func NewTState(thread *Thread) *TState {
	return &TState{
		thread:  thread,
		current: NewIddleState(),
		history: []State{},
	}
}

func (ts *TState) SetState(state State) {
	ts.current = state
	ts.history = append(ts.history, state)
}

func (ts *TState) Do() { ts.current.Do(ts.thread) }
func (ts *TState) Back() {
	if len(ts.history) > 0 {
		ts.current = ts.history[len(ts.history)-1]
		ts.history = ts.history[:len(ts.history)-1]
	}
}

type IddleState struct {
}
type RunningState struct {
}
type DoneState struct {
}
type WaitState struct {
}

func NewIddleState() State   { return &IddleState{} }
func NewWaitState() State    { return &WaitState{} }
func NewRunningState() State { return &RunningState{} }
func NewDoneState() State    { return &DoneState{} }

func (s *IddleState) Next() State   { return NewRunningState() }
func (s *WaitState) Next() State    { return NewRunningState() }
func (s *RunningState) Next() State { return NewDoneState() }
func (s *DoneState) Next() State    { return nil }

func (s *IddleState) Back() State   { return nil }
func (s *WaitState) Back() State    { return NewRunningState() }
func (s *RunningState) Back() State { return NewWaitState() }
func (s *DoneState) Back() State    { return NewRunningState() }

func (s *IddleState) Set(state State) State   { return state }
func (s *DoneState) Set(state State) State    { return state }
func (s *RunningState) Set(state State) State { return state }
func (s *WaitState) Set(state State) State    { return state }

func (s *IddleState) Do(thread *Thread) {
	notify("THREAD_IDDLE", thread.id)

}
func (s *WaitState) Do(thread *Thread) {
	notify("THREAD_WAIT", thread.id)

}
func (s *RunningState) Do(thread *Thread) {
	notify("THREAD_RUNNING", thread.id)

}
func (s *DoneState) Do(thread *Thread) {
	notify("THREAD_DONE", thread.id)
}

func notify(title string, threadID int) {
	c := eapi.Client()
	c.Do(nando.NewRequest(eapi.Notify.String(), &eapi.NotifyRequest{
		Notifier: &events.Notifier{
			Handler: "NOTIFY",
			Event: &events.Event{
				Name:        title,
				Description: fmt.Sprintf("thread_ID: [%d]", threadID),
				Data:        nil,
			},
		},
	}))
}
