package interpreter

import (
	"mary_guica/pkg/tvm/pkg/memory"
)

type Command interface {
	Execute(instruction byte, threadID int, args ...interface{})
}

type base struct {
	memoryManager memory.MemoryManager
	// threadManager *threads.ThreadManager
}

func (b *base) GetMemoryManager() memory.MemoryManager { return b.memoryManager }

// func (b *base) GetThreadManager() *threads.ThreadManager { return b.threadManager }
