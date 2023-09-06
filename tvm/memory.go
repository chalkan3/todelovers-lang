package tvm

type memory struct {
	value []int64
}

func newMemory() *memory {
	return &memory{
		value: make([]int64, 0),
	}
}

func (mem *memory) Get(pc int64) int64 { return mem.value[pc] }
