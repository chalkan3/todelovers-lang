// package main

// import (
// 	"fmt"
// 	"sync"
// )

// type Page struct {
// 	data     []byte
// 	frame    int
// 	dirty    bool
// 	accessed bool
// }

// type FrameAllocator interface {
// 	AllocateFrame() int
// 	FreeFrame(int)
// }

// type MemoryAllocator struct {
// 	frames     []Page
// 	frameMap   map[int]Page
// 	framePool  sync.Pool
// 	frameCount int
// }

// func NewMemoryAllocator(frameCount int) *MemoryAllocator {
// 	return &MemoryAllocator{
// 		frames:     make([]Page, frameCount),
// 		frameMap:   make(map[int]Page),
// 		framePool:  sync.Pool{New: func() interface{} { return new(Page) }},
// 		frameCount: frameCount,
// 	}
// }

// func (ma *MemoryAllocator) AllocatePage() *Page {
// 	frame := ma.framePool.Get().(*Page)
// 	frame.data = make([]byte, 4096)
// 	frame.dirty = false
// 	frame.accessed = false
// 	ma.frames[frame.frame] = *frame
// 	ma.frameMap[frame.frame] = *frame
// 	return frame
// }

// func (ma *MemoryAllocator) FreePage(page *Page) {
// 	ma.frames[page.frame] = Page{}
// 	ma.frameMap[page.frame] = Page{}
// 	ma.framePool.Put(page)
// }

// type MMU struct {
// 	memoryAllocator *MemoryAllocator
// }

// func NewMMU(memoryAllocator *MemoryAllocator) *MMU {
// 	return &MMU{
// 		memoryAllocator: memoryAllocator,
// 	}
// }

// func (mmu *MMU) MapPage(virtualAddress int, physicalAddress int) {
// 	page := mmu.memoryAllocator.AllocatePage()
// 	page.frame = physicalAddress
// 	mmu.memoryAllocator.frames[physicalAddress] = *page
// }

// func (mmu *MMU) UnmapPage(virtualAddress int) {
// 	mmu.memoryAllocator.FreePage(&mmu.memoryAllocator.frames[virtualAddress])
// 	mmu.memoryAllocator.frames[virtualAddress] = Page{}
// }

// func (mmu *MMU) ReadBytes(virtualAddress int, offset ...int) []byte {
// 	page := mmu.memoryAllocator.frames[virtualAddress]
// 	return page.data[offset[0]:offset[1]]
// }

// func (mmu *MMU) WriteByte(virtualAddress int, datapos int, value byte) {
// 	page := mmu.memoryAllocator.frames[virtualAddress]
// 	fmt.Println(virtualAddress % 4096)
// 	page.data[datapos%4096] = value
// 	page.dirty = true
// }

// func main() {
// 	mmu := NewMMU(NewMemoryAllocator(1024))

// 	mmu.MapPage(0x01, 0x01)

// 	str := "Hello, world!"
// 	bytes := []byte(str)

// 	for i, b := range bytes {
// 		mmu.WriteByte(0x01, i, b)
// 	}

// 	fmt.Println(string(mmu.ReadBytes(0x01, 0, len(str))))

// }
package main

import "mary_guica/pkg/tvm"

func main() {
	tvm.T()
}
