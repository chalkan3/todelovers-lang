package tvm

type operands struct {
	value []*register
}

func newOperands() *operands {
	reg1 := newRegister()
	reg2 := newRegister()
	reg3 := newRegister()

	reg1.value = 2
	return &operands{
		value: []*register{
			reg1,
			reg2,
			reg3,
		},
	}
}
