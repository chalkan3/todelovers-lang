package tvm

type Variable struct {
	Name   string
	Value  interface{}
	Adress byte
}

func NewVariable(name string, adress byte, value interface{}) *Variable {
	variable := &Variable{
		Name:   name,
		Adress: adress,
		Value:  value,
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

type Variables map[string]*Variable

func NewVariables() Variables                    { return make(Variables) }
func (vars Variables) Get(name string) *Variable { return vars[name] }
func (vars Variables) NewVariable(params *VariableParams) *Variable {
	newVar := NewVariable(params.Name, params.Adress, params.Value)
	vars[newVar.Name] = newVar

	return newVar
}

type VariableParams struct {
	Name   string
	Adress byte
	Value  interface{}
}
