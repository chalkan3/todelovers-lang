package engine

import "fmt"

type astAssembler struct {
	root        Node
	lexer       *lexer
	nodeFactory *nodeFactory
}

func NewASTAssembler(lexer *lexer, nodeFactory *nodeFactory) *astAssembler {
	return &astAssembler{
		lexer:       lexer,
		nodeFactory: nodeFactory,
	}
}

func (a *astAssembler) parser(token token) Node {
	root := a.nodeFactory.Create(token.Type, token.Value).(*functionCallNode)

	for {
		nextToken := a.lexer.NextToken()

		if isWhiteSpace(nextToken.Type) {
			continue
		}

		if iseof(nextToken.Type) || isCloseParentesis(nextToken.Type) {
			break
		}

		if isNewContext(nextToken.Type) {
			root.Arguments = append(root.Arguments, a.parser(nextToken))
		} else {
			root.Arguments = append(root.Arguments, a.nodeFactory.Create(nextToken.Type, nextToken.Value))
		}

	}

	return root
}

func (a *astAssembler) Assembly(debug bool) *astAssembler {
	root := a.parser(a.lexer.NextToken())
	if debug {
		a.debug(root, "")
	}

	a.root = root

	return a

}

func (a *astAssembler) GetRoot() Node { return a.root }

func (a *astAssembler) debug(node Node, indent string) {
	fmt.Printf("%sType: %s, Token: %v\n", indent, node.Type().String(), node.Token())

	if isNewContext(node.Type()) {
		for _, child := range node.(*functionCallNode).Arguments {
			a.debug(child, indent+"  ")
		}
	}
}
