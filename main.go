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

	vm := tvm.NewTVM()

	code := []byte{0x07, 0x01, 30}
	fmt.Println(string(code))
	// Move o valor de R1 para o registrador R3.

	tvm.LoadCode(vm, code)

}
