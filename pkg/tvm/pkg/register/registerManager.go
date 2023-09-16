package register

const (
	R0 = 0x00
	R1 = 0x01
	R2 = 0x02
)

type RegisterManager interface {
	GetRegister(opID byte) Register
}
type registerManager struct {
	registers []Register
}

func NewRegisterManager() *registerManager {
	return &registerManager{
		registers: []Register{
			NewRegister(),
			NewRegister(),
			NewRegister(),
		},
	}
}

func (op *registerManager) GetRegister(opID byte) Register { return op.registers[opID] }
