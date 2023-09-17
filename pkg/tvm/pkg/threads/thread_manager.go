package threads

import (
	"fmt"
	"mary_guica/pkg/tvm/pkg/events"
)

type ThreadManager interface {
	ThreadPoolSize() int
	GetThread(id int) *Thread
	GetParent(current int) *Thread
	Manage()
	NewThread(id int, parentID int, ec events.EventController) *Thread
}
type threadManager struct {
	pool Pool
}

func NewThreadManager() ThreadManager {
	return &threadManager{
		pool: NewPool(),
	}
}

func (tm *threadManager) NewThread(id int, parentID int, ec events.EventController) *Thread {
	newThread := NewThread(id, parentID, ec)
	tm.pool.Append(newThread)

	return newThread
}

func (tm *threadManager) ThreadPoolSize() int      { return tm.pool.Len() }
func (tm *threadManager) GetThread(id int) *Thread { return tm.pool.Get(id) }
func (tm *threadManager) GetParent(current int) *Thread {
	return tm.pool.Get(tm.pool.Get(current).ParentID())
}

func (tm *threadManager) Manage() {
	for {
		for _, thread := range tm.pool.GetAll() {
			fmt.Println(thread.GetID(), thread.state.String())
		}
	}
}
