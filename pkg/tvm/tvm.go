package tvm

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type types byte

const (
	INT types = iota
	STRING
)

func (t types) Byte() byte {
	return [...]byte{0x00, 0x01}[t]
}

type TVM struct {
	memory      *memory
	heap        *heap
	stack       *stack
	operands    *operands
	pointerMap  *pointerMap
	gc          *garbageCollector
	pc          int
	variables   Variables
	interpreter Interpreter
}

func NewTVM() *TVM {

	tvm := &TVM{
		pc:         0,
		operands:   newOperands(),
		memory:     newMemory(),
		heap:       newHeap(),
		stack:      newStack(),
		pointerMap: newPointerMap(),
		variables:  NewVariables(),
	}

	gc := NewGarbageCollector(tvm)
	tvm.gc = gc

	return tvm
}

func (vm *TVM) AddToHeap(object interface{})                { vm.heap.Add(object) }
func (vm *TVM) GetObjectFromHeap(index int64) interface{}   { return vm.heap.Get(index) }
func (vm *TVM) PushObjectToStack(object interface{})        { vm.stack.Push(object) }
func (vm *TVM) PopObjectFromStack() interface{}             { return vm.stack.Pop() }
func (vm *TVM) MovePC(increment int)                        { vm.pc += increment }
func (vm *TVM) PcPointer(pos int) int                       { return vm.pc + pos }
func (vm *TVM) GetMemoryPos(pos int) byte                   { return vm.memory.Get(pos) }
func (vm *TVM) GetRegister(registerID byte) *register       { return vm.operands.GetRegister(registerID) }
func (vm *TVM) SetMemory(adress int, value byte)            { vm.memory.Add(adress, value) }
func (vm *TVM) CreateVariable(params *VariableParams)       { vm.variables.NewVariable(params) }
func (vm *TVM) GetVariable(name string) *Variable           { return vm.variables.Get(name) }
func (vm *TVM) LenVariables() int                           { return len(vm.variables) }
func (vm *TVM) RegisterInterpreter(interpreter Interpreter) { vm.interpreter = interpreter }

// Carrega um valor de um operando.
func (vm *TVM) LoadOperand(operand *register) {
	operandFromMemory := vm.memory.Get(vm.pc)
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

// func (vm *TVM) ExecuteInstruction(instruction byte) {
// 	// Descodifica a instrução de máquina.
// 	switch instruction {
// 	case 0x01: // Adiciona o operando 1 ao acumulador.
// 		vm.ExecADD(instruction)
// 	case 0x05: // Imprime o valor do acumulador.
// 		vm.ExecPrint(instruction)
// 	case 0x06:
// 		vm.ExecMOV(instruction)
// 	case 0x07:
// 		vm.ExecLOAD(instruction)
// 	case 0x08:
// 		vm.ExecSTORE(instruction)
// 	case 0x09:
// 		vm.Halt()
// 	case 0x0A:
// 		vm.ExecLoadString(instruction)
// 	case 0x0C:
// 		vm.ExecCreateVariable(instruction)
// 	case 0x0D:
// 		vm.ExecLOADFromVariable(instruction)
// 	}
// }

func (vm *TVM) ExecMOV(instruction byte) {

	fromAdress := vm.memory.Get(vm.pc + 1)
	toAdress := vm.memory.Get(vm.pc + 2)

	fromReg := vm.operands.value[fromAdress]
	toReg := vm.operands.value[toAdress]

	toReg.Set(fromReg.value)

	vm.pc += 3
}

func (vm *TVM) LoadCode(code []byte) {
	// Cria uma nova memória virtual para o código do programa.
	programMemory := newMemory()

	// Copia o código do programa para a memória virtual.
	programMemory.Override(code)
	// Executa o código do programa.
	vm.ExecuteCode(programMemory.value)
}

func (vm *TVM) ExecuteCode(code []byte) {
	vm.memory.Override(code)

	for vm.pc < len(vm.memory.value) {
		instruction := vm.memory.value[vm.pc]
		vm.interpreter.Handle(instruction)
	}

}

func (vm *TVM) Prompt() {
	for {
		// Imprime um prompt para o usuário.
		fmt.Println(">> ")

		scanner := bufio.NewScanner(os.Stdin)

		// Lê uma linha do usuário.
		scanner.Scan()

		// Obtém a linha do usuário.
		line := scanner.Text()

		// Se o usuário digitar "exit", saia do loop.
		if line == "exit" {
			break
		}

		// Converte a linha do usuário para um número inteiro.
		instructionNumber, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Erro ao converter a linha do usuário para um número inteiro:", err)
			continue
		}

		// Executa a instrução de máquina.
		vm.interpreter.Handle(byte(instructionNumber))
	}
}

func getIndexOfFirstZero(byteArray []byte) int {
	index := -1

	for i, byte := range byteArray {
		if byte == 0x00 {
			index = i
			break
		}
	}

	return index
}
