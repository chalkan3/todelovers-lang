package engine

import (
	"math"
	"math/rand"
)

type symbolTable struct {
	currentScope *scope
	adress       map[int]bool
}

func NewSymbolTable() *symbolTable {
	globalScope := &scope{
		parent:  nil,
		symbols: make(map[string]Symblo),
	}
	return &symbolTable{
		currentScope: globalScope,
		adress:       make(map[int]bool),
	}
}

func (st *symbolTable) EnterScope() {
	newScope := &scope{
		parent:  st.currentScope,
		symbols: make(map[string]Symblo),
	}
	st.currentScope = newScope
}

func (st *symbolTable) ExitScope() {
	st.currentScope = st.currentScope.parent
}

func (st *symbolTable) AddSymbol(name string, typ Symblo) {
	st.setAdress(&typ)
	st.currentScope.symbols[name] = typ
}

func (st *symbolTable) GetSymbolType(name string) (Symblo, bool) {
	currentScope := st.currentScope
	for currentScope != nil {
		typ, found := currentScope.symbols[name]
		if found {
			return typ, true
		}
		currentScope = currentScope.parent
	}
	return Symblo{}, false
}

func (st *symbolTable) GetSymbolAdress(name string) byte {
	currentScope := st.currentScope
	for currentScope != nil {
		typ, found := currentScope.symbols[name]
		if found {
			return typ.Address
		}
		currentScope = currentScope.parent
	}
	return 0
}

func (st *symbolTable) setAdress(s *Symblo) {
	address := rand.Intn(math.MaxInt64)
	if _, ok := st.adress[address]; !ok {
		st.adress[address] = true

		s.Address = byte(address)
	}

}
func PrintSymbolTable(symbolTable *symbolTable) {
	printScope(symbolTable.currentScope, "")
}
