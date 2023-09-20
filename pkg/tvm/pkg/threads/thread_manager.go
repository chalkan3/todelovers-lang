package threads

import (
	"fmt"
	"mary_guica/pkg/nando"
	eapi "mary_guica/pkg/tvm/internal/api/events"
	"mary_guica/pkg/tvm/pkg/events"
	"time"
)

type ThreadManager interface {
	ThreadPoolSize() int
	GetThread(id int) *Thread
	GetParent(current int) *Thread
	Manage()
	NewThread(id int, parentID int) *Thread
}
type threadManager struct {
	pool Pool
}

func NewThreadManager() ThreadManager {
	return &threadManager{
		pool: NewPool(),
	}
}

func (tm *threadManager) NewThread(id int, parentID int) *Thread {
	newThread := NewThread(id, parentID)
	tm.pool.Append(newThread)

	return newThread
}

func (tm *threadManager) ThreadPoolSize() int { return tm.pool.Len() }
func (tm *threadManager) FreezeThreads(threadID int) {
	for _, t := range tm.pool.GetAll() {
		if t.GetID() != threadID {
			t.SetWait()

		}
	}
}
func (tm *threadManager) ReleaseThreads() {
	for _, t := range tm.pool.GetAll() {
		t.SetWaitRelease()
	}
}

func (tm *threadManager) GetThread(id int) *Thread { return tm.pool.Get(id) }
func (tm *threadManager) GetParent(current int) *Thread {
	return tm.pool.Get(tm.pool.Get(current).ParentID())
}

func (tm *threadManager) Manage() {
	c := &nando.Client{}
	for {
		c.Do(nando.NewRequest(eapi.Notify.String(), &eapi.NotifyRequest{
			Notifier: &events.Notifier{
				Handler: "NOTIFY",
				Event: &events.Event{
					Name:        "THREAD_POOL_SIZE",
					Description: fmt.Sprintf("thread pool size :[%d]", tm.ThreadPoolSize()),
					Data:        nil,
				},
			},
		}))

		time.Sleep(5 * time.Second)

	}
}
