package tvm

type accumulator struct {
	value *register
}

func newAccumulator() *accumulator {
	return &accumulator{
		value: newRegister(),
	}
}
