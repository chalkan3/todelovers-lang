package tvm

type operands struct {
	value []*register
}

func newOperands() *operands {
	return &operands{
		value: make([]*register, 0),
	}
}
