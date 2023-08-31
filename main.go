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
	_ = assembler

}
