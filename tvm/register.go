package tvm

type register struct {
	value byte
}

func newRegister() *register {
	return &register{
		value: 1,
	}
}

func (reg *register) Set(operand byte) { reg.value = operand }
