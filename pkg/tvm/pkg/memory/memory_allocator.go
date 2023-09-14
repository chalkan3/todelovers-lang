package memory

import (
	"sync"
	"unsafe"
)

type MemoryAllocator struct {
	frames     []Page
	frameMap   map[int]Page
	framePool  sync.Pool
	frameCount int
}

func NewMemoryAllocator(frameCount int) *MemoryAllocator {
	return &MemoryAllocator{
		frames:     make([]Page, frameCount),
		frameMap:   make(map[int]Page),
		framePool:  sync.Pool{New: func() interface{} { return new(Page) }},
		frameCount: frameCount,
	}
}

func (ma *MemoryAllocator) AllocatePage(frameID int) *Page {
	frame := ma.framePool.Get().(*Page)
	frame.frame = frameID
	frame.data = make([]byte, 4096)
	frame.dirty = false
	frame.accessed = false
	ma.frames[frame.frame] = *frame
	ma.frameMap[frame.frame] = *frame
	ma.framePool.Put(frame)

	return frame
}

func (ma *MemoryAllocator) FreePage(page *Page) {
	ma.frames[page.frame] = Page{}
	ma.frameMap[page.frame] = Page{}
	ma.framePool.Put(page)
}

func (ma *MemoryAllocator) AllocateStack(size int) int {
	start := ma.framePool.Get().(*Page).frame
	ma.framePool.Put(&ma.frames[start])
	return start

}

func (ma *MemoryAllocator) AllocateHeap(size int) Heap {
	start := ma.framePool.Get().(*Page).frame
	end := start + size - 1
	ma.framePool.Put(&ma.frames[start])
	return Heap{start, end}
}

func (ma *MemoryAllocator) Malloc(heapStart int, size int) unsafe.Pointer {
	for i := heapStart; i <= heapStart+size-1; i++ {
		if !ma.frames[i].dirty {
			ma.frames[i].dirty = true
			return unsafe.Pointer(&ma.frames[i].data[0])
		}
	}

	return nil
}

func (ma *MemoryAllocator) Memcpy(dst unsafe.Pointer, src unsafe.Pointer, n int) {
	if dst == nil || src == nil {
		panic("nil pointer")
	}

	if n < 0 {
		panic("negative length")
	}

	dstBytes := (*[1 << 30]byte)(dst)[:n]
	srcBytes := (*[1 << 30]byte)(src)[:n]

	// Copy the bytes from the source to the destination.
	copy(dstBytes, srcBytes)
}

func (ma *MemoryAllocator) Free(ptr unsafe.Pointer) {
	// Check if the pointer is valid.
	if ptr == nil {
		panic("nil pointer")
	}

	// Get the frame number of the pointer.
	frameNumber := (*Page)(ptr).frame

	// Free the frame from the memory allocator.
	ma.FreePage(&ma.frames[frameNumber])
}
