package tvm

type VM struct {
	registers     [3]int // REG R0, R1, R2
	memory        []int
	allocationMap map[int]*Reference
	stack         []int
	heap          map[int]int
}

func (vm *VM) Interpret(instructions []Instruction) int {
	vm.allocationMap = make(map[int]*Reference)
	vm.stack = []int{}
	vm.heap = make(map[int]int)

	for pc := 0; pc < len(instructions); pc++ {

		instr := instructions[pc]
		switch instr.Opcode {
		case "PUSH":
			dest := instr.Operands[0]
			value := instr.Operands[1]
			ref := &Reference{value: value, refCount: 0, isReachable: false}
			vm.registers[dest] = vm.allocate(ref)

		case "ADD":
			dest := instr.Operands[0]
			src := instr.Operands[1]
			vm.registers[dest] += vm.registers[src]

		case "MOV":
			dest := instr.Operands[0]
			src := instr.Operands[1]
			vm.registers[dest] = vm.registers[src]

		case "STORE":
			src := instr.Operands[0]
			pointer := instr.Operands[1]
			if ref, ok := vm.lookup(src); ok {
				ref.refCount++
				vm.memory[pointer] = ref.value
			} else {
				// Gerenciar erro aqui, o valor não existe.
			}
		case "PUSH_STACK":
			value := instr.Operands[0]
			vm.push(value)
		case "JUMP":
			target := instr.Operands[0]
			pc = target - 1
		case "POP_STACK":
			dest := instr.Operands[0]
			value := vm.pop()
			vm.registers[dest] = value

		}

	}

	vm.markAndSweep()
	vm.gc()

	return vm.registers[0] // Retorne o valor no registrador R0 como resultado.
}

func (vm *VM) gc() {
	// macando os valorer para não alcaçaveis
	for _, ref := range vm.allocationMap {
		ref.isReachable = false
	}

	// marcando os valores que estao a disposicao a partir dos registradores
	for _, reg := range vm.registers {
		if ref, ok := vm.lookup(reg); ok {
			vm.markReachable(ref)
		}
	}

	// free
	for addr, ref := range vm.allocationMap {
		if !ref.isReachable {
			delete(vm.allocationMap, addr)
		}
	}
}

func (vm *VM) markReachable(ref *Reference) {
	if ref.isReachable {
		return
	}
	ref.isReachable = true

	// marcando as referencias para coleta
	if ref.value >= 0 && ref.value < len(vm.memory) {
		if nestedRef, ok := vm.lookup(ref.value); ok {
			vm.markReachable(nestedRef)
		}
	}
}

func (vm *VM) lookup(addr int) (*Reference, bool) {
	ref, ok := vm.allocationMap[addr]
	return ref, ok
}

func (vm *VM) allocate(ref *Reference) int {
	// atribuindo o objeto de referência ao heap.
	addr := len(vm.memory)
	vm.heap[addr] = ref.value
	vm.allocationMap[addr] = ref
	return addr
}

func (vm *VM) push(value int) {
	// push o valor para a pilha.
	vm.stack = append(vm.stack, value)
}

func (vm *VM) pop() int {
	// pop um valor da pilha.
	if len(vm.stack) == 0 {
		// Lidar com erro de pilha vazia.
		return 0
	}
	value := vm.stack[len(vm.stack)-1]
	vm.stack = vm.stack[:len(vm.stack)-1]
	return value
}

func (vm *VM) markAndSweep() {
	// Passo 1: Iniciar a marcação
	for _, ref := range vm.allocationMap {
		ref.isReachable = false
	}

	// Passo 2: Marcar objetos alcançáveis
	for _, reg := range vm.registers {
		if ref, ok := vm.lookup(reg); ok {
			vm.markReachable(ref)
		}
	}

	// Passo 3: Varrer e liberar objetos não alcançáveis
	for addr, ref := range vm.allocationMap {
		if !ref.isReachable {
			// Liberar o objeto não alcançável
			delete(vm.allocationMap, addr)
		}
	}
}
