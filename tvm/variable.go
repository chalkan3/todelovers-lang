package tvm

type Variable struct {
	Name   string
	Value  interface{}
	Adress byte
}

func NewVariable(name string) *Variable {
	variable := &Variable{
		Name: name,
	}

	variable.Value = 0

	return variable
}

func (v *Variable) SetValue(newValue interface{}) {
	v.Value = newValue
}

func (v *Variable) GetValue() interface{} {
	return v.Value
}

func (v *Variable) SetAdress(newValue byte) {
	v.Adress = newValue
}

func (v *Variable) GetAdress() byte {
	return v.Adress
}
