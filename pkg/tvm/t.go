package tvm

import (
	"fmt"
	"mary_guica/pkg/tvm/internal/memory"
	"unsafe"
)

type MyStruct struct {
	Value int
}

func T() {
	manager := memory.NewMemoryManager(memory.NewMemoryAllocator(1024))
	// manager.MapPage(0x00, 0x00)
	manager.AllocateHeap(1024)

	a := "que doideira e essa "
	b := "puts loucura"

	ptr := manager.Malloc(20)
	ptr2 := manager.Malloc(20)

	manager.Memcpy(ptr, unsafe.Pointer(&a), len(a))
	manager.Memcpy(ptr2, unsafe.Pointer(&b), len(b))

	fmt.Println(*(*string)(ptr))
	fmt.Println(*(*string)(ptr2))
	fmt.Println(*(*string)(ptr))
	// manager.Free(ptr)
	// manager.Free(ptr2)

	id := manager.NewStack()

	// Dereference the pointer to get the byte value.

	manager.Push(id, ptr)

	// Converte o array de bytes em um slice de byte
	// Escreve o valor do ponteiro no slice

	// manager.Push(id, 0x03)

	v1 := manager.Pop(id)
	p := unsafe.Pointer(&v1)
	fmt.Println(*(*string)(p))

	// v2 := manager.Pop(id)

	fmt.Println(v1)

}
