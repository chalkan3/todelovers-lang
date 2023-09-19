package todelovers

import (
	"fmt"
	"io/ioutil"
	"mary_guica/pkg/engine"
	"mary_guica/pkg/tvm"
)

type TodeTodeLoversLangTools interface {
	Run(bin string)
	Build(target string)
	GetVM() *tvm.TVM
	BuildAndRun(target string)
}

type tools struct {
	interpreter tvm.Interpreter
	vm          *tvm.TVM
}

func (t *tools) Run(bin string) {

	dataRead, err := ioutil.ReadFile(bin)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// code := []byte{
	// 	interpreter.LOAD, 0x01, tvm.R0,
	// 	interpreter.LOAD, 0x01, tvm.R1,
	// 	interpreter.ADD, tvm.R0, tvm.R2, // r0=2 r1=1 r2=1
	// 	interpreter.PRINT, tvm.R0,
	// 	interpreter.S_THREAD, 0xD,
	// 	interpreter.LOAD, 0x01, tvm.R0,
	// 	interpreter.LOAD, 0x01, tvm.R1,
	// 	interpreter.ADD, tvm.R0, tvm.R2, // r0=2 r1=1 r2=1
	// 	interpreter.PRINT, tvm.R0,
	// 	interpreter.ST_THREAD,
	// 	interpreter.W_THREAD,
	// 	interpreter.HALT,
	// }

	vm := tvm.NewTVM(&tvm.ControlPlaneConfiguration{
		MemoryManager: tvm.MemoryManagerConfig{
			FrameSize: 1024,
		},
		ThreadManager: tvm.ThreadManagerConfig{},
		ProgramManager: tvm.ProgramManagerConfig{
			Code: dataRead,
		},
	})

	vm.ExecuteCode(dataRead)
	// t.vm.LoadCode(code)
}
func (t *tools) Build(target string) {

	dsl, err := engine.File(target)
	if err != nil {
		panic(err)
	}

	symbleTable := engine.NewSymbolTable()
	lexer := engine.NewLexer(dsl).Tokenize()
	nodeFactory := engine.NewNodeFactory()
	assembler := engine.NewASTAssembler(lexer, nodeFactory).Assembly(true)

	root := assembler.GetRoot()
	root.RegisterSymbols(symbleTable, nil)
	engine.PrintSymbolTable(symbleTable)

	code := root.GenerateIntermediateCode(symbleTable)

	filePath := "todbin"

	err = ioutil.WriteFile(filePath, code, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}

func (t *tools) BuildAndRun(target string) {
	dsl, err := engine.File(target)
	if err != nil {
		panic(err)
	}

	symbleTable := engine.NewSymbolTable()
	lexer := engine.NewLexer(dsl).Tokenize()
	nodeFactory := engine.NewNodeFactory()
	assembler := engine.NewASTAssembler(lexer, nodeFactory).Assembly(true)

	root := assembler.GetRoot()
	root.RegisterSymbols(symbleTable, nil)

	code := root.GenerateIntermediateCode(symbleTable)
	vm := tvm.NewTVM(&tvm.ControlPlaneConfiguration{
		MemoryManager: tvm.MemoryManagerConfig{
			FrameSize: 1024,
		},
		ThreadManager: tvm.ThreadManagerConfig{},
		ProgramManager: tvm.ProgramManagerConfig{
			Code: code,
		},
	})
	vm.ExecuteCode(code)

}
func (t *tools) GetVM() *tvm.TVM {
	return t.vm
}
func (t *tools) GetInterpreter() tvm.Interpreter {
	return t.interpreter
}

func New() TodeTodeLoversLangTools {

	return &tools{}
}
