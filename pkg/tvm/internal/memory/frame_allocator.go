package memory

type FrameAllocator interface {
	AllocateFrame() int
	FreeFrame(int)
}
