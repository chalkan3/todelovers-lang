package tvm

type register struct {
	value int64
}

func newRegister() *register {
	return new(register)
}

func (reg *register) Set(operand int64) { reg.value = operand }
