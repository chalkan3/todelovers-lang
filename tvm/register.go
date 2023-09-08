package tvm

type register struct {
	value interface{}
}

func newRegister() *register {
	return &register{
		value: 1,
	}
}

func (reg *register) Set(operand interface{}) { reg.value = operand }
