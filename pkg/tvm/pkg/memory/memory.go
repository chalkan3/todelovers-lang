package memory

type Memory interface {
	Get(pc int) byte
	Memory() []byte
	Set(pc int, value byte)
	Add(adress int, value byte)
	Override(m []byte)
}
type memory struct {
	value []byte
}

func NewMemory() Memory {
	return &memory{
		value: make([]byte, 1024),
	}
}

func (mem *memory) Get(pc int) byte { return mem.value[pc] }
func (mem *memory) Memory() []byte  { return mem.value }

func (mem *memory) Set(pc int, value byte) { mem.value[pc] = value }

func (mem *memory) Add(adress int, value byte) { mem.value[adress] = value }

func (mem *memory) Override(m []byte) {
	n := len(m)
	for i := 0; i < n; i++ {
		mem.value[i] = m[i]
	}
}
