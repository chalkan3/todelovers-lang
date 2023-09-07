package tvm

import "fmt"

type TVM struct {
	memory      *memory
	heap        *heap
	stack       *stack
	operands    *operands
	pointerMap  *pointerMap
	accumulator *accumulator
	gc          *garbageCollector
	pc          int
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
	operandFromMemory := vm.memory.Get(byte(vm.pc))
	operand.Set(operandFromMemory)
	vm.pc++
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

func (vm *TVM) ExecuteInstruction(instruction byte) {
	// Descodifica a instrução de máquina.
	switch instruction {
	case 0x01: // Adiciona o operando 1 ao acumulador.
		vm.accumulator.value.value += vm.operands.value[0].value
	case 0x02: // Subtrai o operando 1 do acumulador.
		vm.accumulator.value.value -= vm.operands.value[0].value
	case 0x03: // Multiplica o acumulador pelo operando 1.
		vm.accumulator.value.value *= vm.operands.value[0].value
	case 0x04: // Divide o acumulador pelo operando 1.
		vm.accumulator.value.value /= vm.operands.value[0].value
	case 0x05: // Imprime o valor do acumulador.
		fmt.Println(vm.accumulator.value.value)
	case 0x06:
		vm.ExecMOV(instruction)
	case 0x07:
		vm.ExecLOAD(instruction)
	default:
		fmt.Println("Instrução de máquina desconhecida.")
	}
}

func (vm *TVM) ExecMOV(instruction byte) {
	// Obtém o operando 1.
	operand1 := vm.operands.value[0]

	// Obtém o registrador de destino.
	register := byte(instruction & 0x02)
	// Move o valor do operando 1 para o registrador de destino.
	vm.operands.value[register] = operand1
}

func LoadCode(vm *TVM, code []byte) {
	// Cria uma nova memória virtual para o código do programa.
	programMemory := newMemory()

	// Copia o código do programa para a memória virtual.
	programMemory.Override(code)
	// Executa o código do programa.
	vm.ExecuteCode(programMemory.value)
}

func (vm *TVM) ExecuteCode(code []byte) {
	vm.memory.Override(code)
	// Apontador para o próximo byte de código a ser executado.
	pc := 0

	// Executa o código até que ele chegue ao fim.
	for pc < len(vm.memory.value) {
		// Obtém a próxima instrução de máquina.
		instruction := vm.memory.value[pc]

		// Executa a instrução de máquina.
		vm.ExecuteInstruction(instruction)

		// Avança para o próximo byte de código.
		pc++
	}

	fmt.Println(vm.operands.value[0].value)
}

func (vm *TVM) ExecLOAD(instruction byte) {
	// Obtém o operando 1.
	operand1 := vm.operands.value[0]

	// Obtém o registrador de destino.
	register := byte(instruction & 0x02)

	// Carrega o valor da memória no registrador.
	vm.operands.value[register].value = vm.memory.value[operand1.value]
}
