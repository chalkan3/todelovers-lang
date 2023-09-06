package main

import (
	"fmt"
	"mary_guica/engine"
	"mary_guica/tvm"
)

// Arrumar a arvore sintaxica
// Realizar a criação do codigo intermediario
// fazer a  geração de código objeto com assembly

func main() {
	dsl, err := engine.File("main.todelovers")
	if err != nil {
		panic(err)
	}

	symbleTable := engine.NewSymbolTable()
	lexer := engine.NewLexer(dsl).Tokenize()
	nodeFactory := engine.NewNodeFactory()
	assembler := engine.NewASTAssembler(lexer, nodeFactory).Assembly(false)
	// logger := engine.NewLogger(&engine.LoggerConfig{
	// 	Enable:     true,
	// 	Mode:       engine.Stack,
	// 	BufferSize: 100,
	// })

	root := assembler.GetRoot()
	root.RegisterSymbols(symbleTable, nil)

	fmt.Println(root.GenerateIntermediateCode())

	instructions := []tvm.Instruction{
		{Opcode: "PUSH", Operands: [2]int{0, 10}},
		{Opcode: "PUSH", Operands: [2]int{1, 20}},
		{Opcode: "ADD", Operands: [2]int{0, 1}},
	}

	vm := tvm.VM{}
	result := vm.Interpret(instructions)
	fmt.Println(result)

}
