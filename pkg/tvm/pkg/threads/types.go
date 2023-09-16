package threads

import (
	"mary_guica/pkg/tvm/pkg/memory"
	"mary_guica/pkg/tvm/pkg/register"
)

type (
	Register = register.Register
	Memory   = memory.Memory
)

var (
	NewMemory = memory.NewMemory
)
