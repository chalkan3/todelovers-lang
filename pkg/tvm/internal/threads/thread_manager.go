package threads

import (
	"fmt"
	"mary_guica/pkg/interpreter"
)

type ThreadManager struct {
	threads threadPool
	interpreter Interpreter
}

func NewThreadManager() *ThreadManager {
	return &ThreadManager{
		threads: []*Thread{},
		interpreter: New
	}
}

func (tm *ThreadManager) NewThread(interpreter Interpreter, parentID int) *Thread {
	newThread := NewThread(len(tm.threads), parentID, interpreter)
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
