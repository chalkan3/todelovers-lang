package tvm

type register struct {
	value interface{}
}

func newRegister() *register {
	return &register{
		value: 1,
	}
}

func (reg *register) Set(operand interface{})           { reg.value = operand }
func (reg *register) Get(operand interface{}) *register { return reg }
func (reg *register) Value() interface{}                { return reg.value }
