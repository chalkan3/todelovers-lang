package memory

import "unsafe"

type Stack struct {
	allocator *MemoryAllocator
	top       unsafe.Pointer
}

func NewStack(allocator *MemoryAllocator) *Stack {
	return &Stack{
		allocator: allocator,
		top:       nil,
	}
}

func (stack *Stack) Push(value unsafe.Pointer) {
	// Allocate a new page from the frame pool if necessary.
	page := stack.allocator.AllocatePage()

	if stack.top == nil {

		stack.top = unsafe.Pointer(&page.data[0])
	} else {
		// Update the pointer to the next page in the stack.
		(*Page)(stack.top).next = page
		stack.top = unsafe.Pointer(&page.data[0])
	}

	// Write the value to the top of the stack.
	*(*unsafe.Pointer)(stack.top) = value
}

func (stack *Stack) Pop() unsafe.Pointer {
	// Check if the stack is empty.
	if stack.top == nil {
		panic("empty stack")
	}

	// Get the value from the top of the stack.
	value := *(*unsafe.Pointer)(stack.top)

	// Update the pointer to the next page in the stack.
	currentPage := (*Page)(stack.top)
	stack.top = unsafe.Pointer(currentPage.next)

	// Return the page to the frame pool.
	stack.allocator.framePool.Put(currentPage)

	return value
}

func (stack *Stack) IsEmpty() bool {
	return stack.top == nil
}
