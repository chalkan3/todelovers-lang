package tvm

type ThreadManager struct {
	threads []*Thread
}

func NewThreadManager() *ThreadManager {
	return &ThreadManager{
		threads: []*Thread{},
	}
}

func (tm *ThreadManager) NewThread(interpreter Interpreter) *Thread {
	newThread := NewThread(len(tm.threads), interpreter)
	tm.threads = append(tm.threads, newThread)

	return newThread
}

func (tm *ThreadManager) GetThread(id int) *Thread { return tm.threads[id] }
