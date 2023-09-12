package register

type Register interface {
	Set(operand interface{})
	Get(operand interface{}) Register
	Value() interface{}
}
type register struct {
	value interface{}
}

func NewRegister() Register {
	return &register{
		value: 1,
	}
}

func (reg *register) Set(operand interface{})          { reg.value = operand }
func (reg *register) Get(operand interface{}) Register { return reg }
func (reg *register) Value() interface{}               { return reg.value }
