package program

type Fork interface {
	New(key byte, program Program)
	GetProgram(key byte) Program
}
type fork map[byte]Program

func NewFork() Fork {
	return make(fork)
}

func (f fork) New(key byte, program Program) { f[key] = program }
func (f fork) GetProgram(key byte) Program   { return f[key] }
