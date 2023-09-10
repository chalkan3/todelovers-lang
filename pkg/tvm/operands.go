package tvm

const (
	R0 = 0x00
	R1 = 0x01
	R2 = 0x02
)

type operands struct {
	value []*register
}

func newOperands() *operands {
	return &operands{
		value: []*register{
			newRegister(),
			newRegister(),
			newRegister(),
		},
	}
}

func (op *operands) GetRegister(opID byte) *register { return op.value[opID] }
