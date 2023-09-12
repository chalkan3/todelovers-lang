package threads

type Pool interface {
	Get(id int) Thread
}
type pool []*Thread

func NewPool() Pool {
	return pool{}
}

func (p *pool) Get(id int) Thread {

}
