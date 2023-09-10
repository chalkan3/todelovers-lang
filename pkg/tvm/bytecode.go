package tvm

type bytecode struct {
	value []int64
}

func NewByteCode() *bytecode {
	return &bytecode{
		value: make([]int64, 0),
	}
}
