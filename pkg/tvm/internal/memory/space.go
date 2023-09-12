package memory

type Space interface {
	CreateSpace(spaceName string, memory Memory)
	GetSpace(spaceName string) Memory
}

type space struct {
	spaces map[string]Memory
}

func (s *space) CreateSpace(spaceName string, memory Memory) {
	s.spaces[spaceName] = memory
}
func (s *space) GetSpace(spaceName string) Memory {
	return s.spaces[spaceName]
}

func NewSpace() *space {
	return &space{
		spaces: map[string]Memory{},
	}
}
