package engine

import "fmt"

type scope struct {
	parent      *scope
	symbols     map[string]Symblo
	childScopes []*scope
	name        string
}

func NewScope(name string) *scope {
	return &scope{
		name:        name,
		parent:      nil,
		symbols:     make(map[string]Symblo),
		childScopes: []*scope{},
	}
}

// AddChildScope adds a child scope to a parent scope
func (s *scope) AddChildScope(childScope *scope) {
	s.childScopes = append(s.childScopes, childScope)
	childScope.parent = s
}

func printScope(scope *scope, indent string) {
	fmt.Printf("%sEscopo: %s\n", indent, scope.name)
	for symbol, typ := range scope.symbols {
		fmt.Printf("%s%s: %s\n", indent, symbol, typ)
	}

	for _, childScope := range scope.childScopes {
		printScope(childScope, indent+"  ")
	}
}
