package register

const (
	R0 = 0x00
	R1 = 0x01
	R2 = 0x02
)

type Operands interface {
	GetRegister(opID byte) Register
}
type operands struct {
	value []Register
}

func NewOperands() *operands {
	return &operands{
		value: []Register{
			NewRegister(),
			NewRegister(),
			NewRegister(),
		},
	}
}

func (op *operands) GetRegister(opID byte) Register { return op.value[opID] }
