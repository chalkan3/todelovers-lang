package threads

import (
	"mary_guica/pkg/interpreter"
	"mary_guica/pkg/tvm/pkg/memory"
	"mary_guica/pkg/tvm/pkg/register"
)

type (
	Interpreter = interpreter.Interpreter
	Operands    = register.Operands
	Register    = register.Register
	Memory      = memory.Memory
)

var (
	NewOperands = register.NewOperands
	NewMemory   = memory.NewMemory
)
