package main

import (
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

	code := root.GenerateIntermediateCode()

	vm := tvm.NewTVM()

	// Initialize a byte slice for bytecode

	// code := []byte{
	// 	0x07, 0x04, 0x00, // LOAD 1 R0
	// 	0x07, 0x04, 0x01, // LOAD 1 R1
	// 	0x0A, 0x02, 0x02, 'h', 'i', // LOAD "HI" R0
	// 	0x08, 100, 0x02,
	// 	0x0B, 100, 0x02,
	// 	0x09,
	// }

	// var pos int = 100
	// for _, char := range str {
	// 	// Convert the character to its ASCII value and add it to the bytecode
	// 	code = append(code, 0x07)
	// 	code = append(code, byte(char))
	// 	code = append(code, 0x00)

	// 	code = append(code, 0x08)
	// 	code = append(code, byte(pos))
	// 	code = append(code, 0x00)

	// 	pos++
	// }

	// code = append(code, 0x09)

	// var hundred byte
	// hundred = 0x30

	// for _, s := range a {
	// 	code = append(code, 0x07)
	// 	code = append(code, s)
	// 	code = append(code, 0x00)

	// 	code = append(code, 0x08)
	// 	code = append(code, hundred)
	// 	code = append(code, 0x00)

	// 	hundred = hundred + 1

	// }

	tvm.LoadCode(vm, code)

}
