package tvm

type TVM struct {
	memory      *memory
	heap        *heap
	stack       *stack
	operands    *operands
	pointerMap  *pointerMap
	accumulator *accumulator
	gc          *garbageCollector
	pc          int64
}

func NewTVM() *TVM {

	tvm := &TVM{
		accumulator: newAccumulator(),
		operands:    newOperands(),
		memory:      newMemory(),
		heap:        newHeap(),
		stack:       newStack(),
		pointerMap:  newPointerMap(),
	}

	gc := NewGarbageCollector(tvm)
	tvm.gc = gc

	return tvm
}

func (vm *TVM) AddToHeap(object interface{})              { vm.heap.Add(object) }
func (vm *TVM) GetObjectFromHeap(index int64) interface{} { return vm.heap.Get(index) }
func (vm *TVM) PushObjectToStack(object interface{})      { vm.stack.Push(object) }
func (vm *TVM) PopObjectFromStack() interface{}           { return vm.stack.Pop() }

// Carrega um valor de um operando.
func (vm *TVM) LoadOperand(operand *register) {
	operandFromMemory := vm.memory.Get(vm.pc)
	operand.Set(operandFromMemory)
}

func (vm *TVM) CreatePointer(address int64) *pointer {
	pointer := newPointer(address)

	vm.pointerMap.Set(address, pointer)

	return pointer
}

func (vm *TVM) GetObjectFromPointer(pointer *pointer) interface{} {
	return vm.heap.GetObjectFromPointer(pointer)
}

func (vm *TVM) SetObjectToPointer(pointer *pointer, object interface{}) {
	vm.heap.SetObjectToPointer(pointer, object)
}

func (vm *TVM) GetObjectReferences(object interface{}) []int64 {
	// Create a new list to store the object references.
	references := make([]int64, 0)

	// Iterate over the heap.
	for i, heapObject := range vm.heap.GetAll() {
		// If the heap object is the same as the specified object.
		if heapObject == object {
			// Add the index of the heap object to the list of references.
			references = append(references, int64(i))
		}

		// If the heap object is a pointer.
		if pointer, ok := vm.pointerMap.Get(int64(i)); ok {
			// Get the object references from the heap object.
			references = append(references, vm.GetObjectReferences(pointer.GetAddress())...)
		}
	}

	// Return the list of object references.
	return references
}
