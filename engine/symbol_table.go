package engine

type SymbolTable struct {
	Functions map[string]*FunctionNode
	Variables map[string]interface{}
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		Functions: make(map[string]*FunctionNode),
		Variables: make(map[string]interface{}),
	}
}

func (s *SymbolTable) AddFunction(name string, function *FunctionNode) {
	s.Functions[name] = function
}

func (s *SymbolTable) GetFunction(name string) *FunctionNode {
	return s.Functions[name]
}
