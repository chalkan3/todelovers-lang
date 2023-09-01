package engine

type symbolTable struct {
	Functions map[string]*functionNode
	Variables map[string]interface{}
}

func NewSymbolTable() *symbolTable {
	return &symbolTable{
		Functions: make(map[string]*functionNode),
		Variables: make(map[string]interface{}),
	}
}

func (s *symbolTable) AddFunction(name string, function *functionNode) {
	s.Functions[name] = function
}

func (s *symbolTable) GetFunction(name string) *functionNode {
	return s.Functions[name]
}
