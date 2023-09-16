package program

type ProgramManager interface {
	NewFork(id byte, code []byte)
	Next(forkID ...byte)
	Jump(pos int, forkID ...byte)
	Current(forkID ...byte) int
	Code(forkID ...byte) []byte
	Instruction(forkID ...byte) byte
	GetAdressValue(pos int, forkID ...byte) byte
}
type programManager struct {
	main  Program
	forks Fork
}

func NewProgramManager(code []byte) ProgramManager {
	return &programManager{
		main:  NewProgram(code),
		forks: NewFork(),
	}
}

func (p *programManager) getMainProgram() Program        { return p.main }
func (p *programManager) getForkProgram(id byte) Program { return p.forks.GetProgram(id) }
func (p *programManager) setProgram(id ...byte) Program {
	handler := p.getMainProgram()
	if len(id) > 0 {
		handler = p.getForkProgram(id[0])
	}

	return handler

}

func (p *programManager) NewFork(id byte, code []byte) { p.forks.New(id, NewProgram(code)) }
func (p *programManager) Next(forkID ...byte)          { p.setProgram(forkID...).Next() }
func (p *programManager) Jump(pos int, forkID ...byte) { p.setProgram(forkID...).Jump(pos) }
func (p *programManager) Current(forkID ...byte) int   { return p.setProgram().Current() }
func (p *programManager) Code(forkID ...byte) []byte   { return p.setProgram(forkID...).Code() }
func (p *programManager) Instruction(forkID ...byte) byte {
	return p.setProgram(forkID...).Instruction()
}

func (p *programManager) GetAdressValue(pos int, forkID ...byte) byte {
	return p.setProgram(forkID...).GetAdressValue(pos)
}
