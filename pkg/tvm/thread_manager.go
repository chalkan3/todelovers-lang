package tvm

import "fmt"

type ThreadManager struct {
	threads []*Thread
}

func NewThreadManager() *ThreadManager {
	return &ThreadManager{
		threads: []*Thread{},
	}
}

func (tm *ThreadManager) NewThread(parentID int) *Thread {
	newThread := NewThread(len(tm.threads), parentID)
	tm.threads = append(tm.threads, newThread)

	return newThread
}

func (tm *ThreadManager) ThreadPoolSize() int { return len(tm.threads) }

func (tm *ThreadManager) GetThread(id int) *Thread { return tm.threads[id] }
func (tm *ThreadManager) GetParent(current int) *Thread {
	return tm.GetThread(tm.threads[current].parentID)
}

func (tm *ThreadManager) Manage() {
	for {
		for _, thread := range tm.threads {
			fmt.Println(thread.GetID(), thread.state.String())
		}
	}
}
