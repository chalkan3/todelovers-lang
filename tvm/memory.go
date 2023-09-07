package tvm

type memory struct {
	value []byte
}

func newMemory() *memory {
	return &memory{
		value: []byte{},
	}
}

func (mem *memory) Get(pc byte) byte { return mem.value[pc] }
func (mem *memory) Override(m []byte) {
	mem.value = m
}
