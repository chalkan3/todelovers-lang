package tvm

type memory struct {
	value []byte
}

func newMemory() *memory {
	return &memory{
		value: make([]byte, 1024),
	}
}

func (mem *memory) Get(pc int) byte        { return mem.value[pc] }
func (mem *memory) Set(pc int, value byte) { mem.value[pc] = value }

func (mem *memory) Add(adress int, value byte) { mem.value[adress] = value }

func (mem *memory) Override(m []byte) {
	n := len(m)

	for i := 0; i < n; i++ {
		mem.value[i] = m[i]
	}
}
