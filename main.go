package main

import (
	"mary_guica/engine"
)

func main() {
	dsl, err := engine.File("main.todelovers")
	if err != nil {
		panic(err)
	}

	lexer := engine.NewLexer(dsl).Tokenize()
	nodeFactory := engine.NewNodeFactory()
	assembler := engine.NewASTAssembler(lexer, nodeFactory).Assembly(true)
	logger := engine.NewLogger(&engine.LoggerConfig{
		Enable:     true,
		Mode:       engine.Stack,
		BufferSize: 100,
	})

	logger.Log(assembler.GetRoot())

}
