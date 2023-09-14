package memory

import (
	"fmt"
	"sync"
	"unsafe"
)

type Stack struct {
	ma   *MemoryAllocator
	top  int
	lock sync.Mutex
}

func NewStack(ma *MemoryAllocator) *Stack {
	return &Stack{
		ma:   ma,
		top:  -1,
		lock: sync.Mutex{},
	}
}

func (s *Stack) Push(data []byte) {
	s.lock.Lock()
	defer s.lock.Unlock()

	// Aloca uma página para armazenar os dados da pilha
	page := s.ma.AllocatePage(s.top + 1)

	// Copia os dados para a página alocada
	copy(page.data, data)
	s.ma.frames[s.top+1].dirty = true

	// Atualiza o topo da pilha
	s.top++
}

func (s *Stack) Pop() ([]byte, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.top == -1 {
		return nil, fmt.Errorf("a pilha está vazia")
	}

	// Copia os dados do topo da pilha
	data := make([]byte, len(s.ma.frames[s.top].data))
	s.ma.Memcpy(unsafe.Pointer(&data[0]), unsafe.Pointer(&s.ma.frames[s.top].data[0]), len(data))

	// Libera a página do topo da pilha
	s.ma.FreePage(&s.ma.frames[s.top])

	// Atualiza o topo da pilha
	s.top--

	return data, nil
}
