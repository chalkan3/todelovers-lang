package program

type Program interface {
	Next()
	Jump(pos int)
	GetAdressValue(pos int) byte
	Current() int
	Code() []byte
	Instruction() byte
}
type program struct {
	pointer int
	code    []byte
}

func NewProgram(code []byte) Program {
	return &program{
		pointer: 0,
		code:    code,
	}
}

func (p *program) Next()                       { p.pointer++ }
func (p *program) Jump(pos int)                { p.pointer += pos }
func (p *program) Current() int                { return p.pointer }
func (p *program) Code() []byte                { return p.code }
func (p *program) Instruction() byte           { return p.code[p.pointer] }
func (p *program) GetAdressValue(pos int) byte { return p.code[p.pointer+pos] }
