package engine

type symbolTable struct {
	currentScope *scope
}

func NewSymbolTable() *symbolTable {
	globalScope := &scope{
		parent:  nil,
		symbols: make(map[string]tokenType),
	}
	return &symbolTable{
		currentScope: globalScope,
	}
}

func (st *symbolTable) EnterScope() {
	newScope := &scope{
		parent:  st.currentScope,
		symbols: make(map[string]tokenType),
	}
	st.currentScope = newScope
}

func (st *symbolTable) ExitScope() {
	st.currentScope = st.currentScope.parent
}

func (st *symbolTable) AddSymbol(name string, typ tokenType) {
	st.currentScope.symbols[name] = typ
}

func (st *symbolTable) GetSymbolType(name string) (tokenType, bool) {
	currentScope := st.currentScope
	for currentScope != nil {
		typ, found := currentScope.symbols[name]
		if found {
			return typ, true
		}
		currentScope = currentScope.parent
	}
	return eof, false
}

func PrintSymbolTable(symbolTable *symbolTable) {
	printScope(symbolTable.currentScope, "")
}
