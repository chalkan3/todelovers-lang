package threads

import (
	"mary_guica/pkg/interpreter"
	"mary_guica/pkg/tvm/internal/memory"
	"mary_guica/pkg/tvm/internal/register"
)

type (
	Interpreter = interpreter.Interpreter
	Operands    = register.Operands
	Register    = register.Register
	Memory      = memory.Memory
)

var (
	NewInterpreter = interpreter.NewInterpreter
	NewOperands    = register.NewOperands
	NewMemory      = memory.NewMemory
)
