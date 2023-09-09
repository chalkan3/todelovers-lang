package tvm

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TVM struct {
	memory     *memory
	heap       *heap
	stack      *stack
	operands   *operands
	pointerMap *pointerMap
	gc         *garbageCollector
	pc         int
	variables  map[string]*Variable
}

func NewTVM() *TVM {

	tvm := &TVM{
		pc:         0,
		operands:   newOperands(),
		memory:     newMemory(),
		heap:       newHeap(),
		stack:      newStack(),
		pointerMap: newPointerMap(),
		variables:  make(map[string]*Variable),
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
		vm.ExecADD(instruction)
	case 0x05: // Imprime o valor do acumulador.
		vm.ExecPrint(instruction)
	case 0x06:
		vm.ExecMOV(instruction)
	case 0x07:
		vm.ExecLOAD(instruction)
	case 0x08:
		vm.ExecSTORE(instruction)
	case 0x09:
		vm.Halt()
	case 0x0A:
		vm.ExecLoadString(instruction)
	case 0x0C:
		vm.ExecCreateVariable(instruction)
	case 0x0D:
		vm.ExecLOADFromVariable(instruction)
	}
}

func (vm *TVM) ExecMOV(instruction byte) {

	fromAdress := vm.memory.Get(vm.pc + 1)
	toAdress := vm.memory.Get(vm.pc + 2)

	fromReg := vm.operands.value[fromAdress]
	toReg := vm.operands.value[toAdress]

	toReg.Set(fromReg.value)

	vm.pc += 3
}

func LoadCode(vm *TVM, code []byte) {
	// Cria uma nova memória virtual para o código do programa.
	programMemory := newMemory()

	// Copia o código do programa para a memória virtual.
	programMemory.Override(code)
	// Executa o código do programa.
	vm.ExecuteCode(programMemory.value)
}

func (vm *TVM) ExecADD(instruction byte) {
	fromAdress := vm.memory.Get(vm.pc + 1)
	toAdress := vm.memory.Get(vm.pc + 2)

	reg1 := vm.operands.value[fromAdress]
	reg2 := vm.operands.value[toAdress]

	reg1.value = reg1.value.(int) + reg2.value.(int)
	vm.pc += 3
}

func (vm *TVM) ExecPrint(instruction byte) {
	registerAddress := vm.memory.Get(vm.pc + 1)
	reg1 := vm.operands.value[registerAddress]

	fmt.Println(reg1.value)

	vm.pc += 2
}

func (vm *TVM) ExecSTORE(instruction byte) {
	registerAdress := vm.memory.Get(vm.pc + 1)
	saveToAdress := vm.memory.Get(vm.pc + 2)

	register := vm.operands.value[registerAdress]
	isSTR, ok := register.value.(string)
	if ok {
		for _, c := range isSTR {
			vm.memory.Add(int(saveToAdress), byte(c))
			saveToAdress++
		}
	} else {
		vm.memory.Add(int(saveToAdress), register.value.(byte))
	}

	// Increment the program counter.
	vm.pc += 3
}

func (vm *TVM) Halt() {
	// fmt.Println("reg:1", vm.operands.value[0].value)
	// fmt.Println("reg:2", vm.operands.value[1].value)
	// fmt.Println("reg:3", vm.operands.value[2].value)
	os.Exit(0)
}

func (vm *TVM) ExecuteCode(code []byte) {
	vm.memory.Override(code)
	// Apontador para o próximo byte de código a ser executado.

	// Executa o código até que ele chegue ao fim.
	// fmt.Println("TVM MEMORY SETTED", len(vm.memory.value))
	for vm.pc < len(vm.memory.value) {
		// Obtém a próxima instrução de máquina.
		instruction := vm.memory.value[vm.pc]
		// fmt.Printf("pc = %d\t0x%02x\n", vm.pc, instruction)
		// Executa a instrução de máquina.
		vm.ExecuteInstruction(instruction)
	}

}

func (vm *TVM) ExecLOAD(instruction byte) {
	address := vm.memory.value[vm.pc+1]
	register := vm.memory.value[vm.pc+2]

	vm.operands.value[register].Set(address)
	vm.pc += 3
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
		vm.ExecuteInstruction(byte(instructionNumber))
	}
}

func (vm *TVM) ExecLoadString(instruction byte) {
	regNumber := vm.memory.value[vm.pc+1]

	register := vm.operands.value[regNumber]

	// Get the length of the string from the instruction.
	length := vm.memory.value[vm.pc+2]

	// Create a byte slice to store the string data.
	strData := make([]byte, length)

	// Copy the string data from the instruction.
	for i := byte(0); i < length; i++ {
		n := byte(vm.pc+3) + i
		strData[i] = vm.memory.value[n]
	}

	// Convert the byte slice to a string.
	str := string(strData)

	// Set the string to a register (assuming you have a register available).
	register.Set(str)

	vm.pc += 3 + int(length)
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

func (vm *TVM) ExecCreateVariable(instruction byte) {
	// Get the variable name length from the instruction.
	nameLength := vm.memory.value[vm.pc+1]

	// Create a byte slice to store the variable name.
	nameBytes := make([]byte, nameLength)

	// Copy the variable name from memory.
	for i := byte(0); i < nameLength; i++ {
		nameBytes[i] = vm.memory.value[byte(vm.pc+2)+i]
	}

	// Convert the byte slice to a string (variable name).
	variableName := string(nameBytes)

	vm.memory.value[0xC8] = 0

	newVariable := NewVariable(variableName)

	sizeVariables := len(vm.variables)
	newVariable.SetAdress(0xC8 + byte(sizeVariables))

	vm.variables[variableName] = newVariable

	vm.pc += int(nameLength) + 2 // Adjust for the variable name length and opcode.
}

func (vm *TVM) ExecLOADFromVariable(instruction byte) {
	// Get the memory address from the instruction (assuming it's a 1-byte instruction).
	register := vm.memory.value[vm.pc+1]

	nameLength := vm.memory.value[vm.pc+2]

	// Create a byte slice to store the variable name.
	nameBytes := make([]byte, nameLength)

	// Copy the variable name from memory.
	for i := byte(0); i < nameLength; i++ {
		nameBytes[i] = vm.memory.value[byte(vm.pc+3)+i]
	}

	// Convert the byte slice to a string (variable name).
	variableName := string(nameBytes)

	address := vm.variables[variableName].GetAdress()

	// Get the destination register from the instruction (assuming it's a 1-byte instruction).
	memoryValue := vm.memory.value[address]
	vm.operands.value[register].Set(memoryValue)
	vm.pc += int(nameLength) + 3
}
