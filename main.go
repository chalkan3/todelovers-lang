package main

import (
	"fmt"
	"mary_guica/engine"
)

// Arrumar a arvore sintaxica
// Realizar a criação do codigo intermediario
// fazer a  geração de código objeto com assembly

func main() {
	dsl, err := engine.File("main.todelovers")
	if err != nil {
		panic(err)
	}

	lexer := engine.NewLexer(dsl).Tokenize()
	nodeFactory := engine.NewNodeFactory()
	assembler := engine.NewASTAssembler(lexer, nodeFactory).Assembly(true)
	// logger := engine.NewLogger(&engine.LoggerConfig{
	// 	Enable:     true,
	// 	Mode:       engine.Stack,
	// 	BufferSize: 100,
	// })

	root := assembler.GetRoot()
	fmt.Println(root.GenerateIntermediateCode())

}
