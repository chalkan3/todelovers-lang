package threads

type Pool interface {
	Get(id int) *Thread
	Append(thread *Thread)
	GetAll() []*Thread
	Len() int
}
type pool struct {
	threads []*Thread
}

func NewPool() Pool {
	return &pool{
		threads: []*Thread{},
	}
}

func (p *pool) Get(id int) *Thread { return p.threads[id] }
func (p *pool) GetAll() []*Thread  { return p.threads }
func (p *pool) Append(thread *Thread) {
	p.threads = append(p.threads, thread)
}
func (p *pool) Len() int { return len(p.threads) }
