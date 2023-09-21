package memory

import (
	"unsafe"
)

type MemoryManager interface {
	MapPage(virtualAddress int, physicalAddress int)
	UnmapPage(virtualAddress int)
	ReadBytes(virtualAddress int, offset ...int) []byte
	WriteByte(virtualAddress int, datapos int, value byte)
	Malloc(size int) unsafe.Pointer
	Memcpy(dst unsafe.Pointer, src unsafe.Pointer, n int)
	Free(dst unsafe.Pointer)
	MapHeap(heap Heap)
	AllocateHeap(size int)
	AllocateStack(size int)
	Stack() *Stack
	APIPath() string
}

type memoryManager struct {
	memoryAllocator *MemoryAllocator
	heap            Heap
	stack           *Stack
}

func NewMemoryManager(memoryAllocator *MemoryAllocator) MemoryManager {
	return &memoryManager{
		memoryAllocator: memoryAllocator,
		stack:           NewStack(memoryAllocator),
	}
}

func (mmu *memoryManager) Malloc(size int) unsafe.Pointer {
	return mmu.memoryAllocator.Malloc(mmu.heap.Start, size)
}

func (mmu *memoryManager) Memcpy(dst unsafe.Pointer, src unsafe.Pointer, n int) {
	mmu.memoryAllocator.Memcpy(dst, src, n)
}

func (mmu *memoryManager) Free(dst unsafe.Pointer) {
	mmu.memoryAllocator.Free(dst)
}

func (mmu *memoryManager) MapPage(virtualAddress int, physicalAddress int) {
	page := mmu.memoryAllocator.AllocatePage(physicalAddress)
	mmu.memoryAllocator.frames[physicalAddress] = *page
}

func (mmu *memoryManager) UnmapPage(virtualAddress int) {
	mmu.memoryAllocator.FreePage(&mmu.memoryAllocator.frames[virtualAddress])
	mmu.memoryAllocator.frames[virtualAddress] = Page{}
}

func (mmu *memoryManager) ReadBytes(virtualAddress int, offset ...int) []byte {
	page := mmu.memoryAllocator.frames[virtualAddress]
	return page.data[offset[0]:offset[1]]
}

func (mmu *memoryManager) WriteByte(virtualAddress int, datapos int, value byte) {
	page := mmu.memoryAllocator.frames[virtualAddress]
	page.data[datapos%4096] = value
	page.dirty = true
}

func (mmu *memoryManager) MapHeap(heap Heap) {
	for i := heap.Start; i <= heap.End; i++ {
		mmu.MapPage(i, i)
	}
}

func (mmu *memoryManager) AllocateHeap(size int) {
	heap := mmu.memoryAllocator.AllocateHeap(size)
	mmu.heap = heap
}

func (mmu *memoryManager) AllocateStack(size int) {
	mmu.memoryAllocator.AllocateStack(size)
}

func (mmu *memoryManager) APIPath() string { return "memory.manager" }

func (mmu *memoryManager) Stack() *Stack {
	return mmu.stack
}
