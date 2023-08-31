package engine

import "fmt"

type parser func(tokens []Token) Node
type ASTAssembler struct {
	root        Node
	lexer       *Lexer
	nodeFactory *NodeFactory
}

func NewASTAssembler(lexer *Lexer, nodeFactory *NodeFactory) *ASTAssembler {
	return &ASTAssembler{
		lexer:       lexer,
		nodeFactory: nodeFactory,
	}
}

func (a *ASTAssembler) parser(token Token) Node {

	root := a.nodeFactory.Create(token.Type, token.Value).(*FunctionCallNode)

	for {
		nextToken := a.lexer.NextToken()

		if nextToken.Type == EOF {
			break
		}
		if nextToken.Type == WHITESPACE {
			continue
		}
		if nextToken.Type == CLOSE_PAREN {
			break
		}
		if nextToken.Type == OPEN_PAREN || nextToken.Type == CONTEXT || nextToken.Type == DEF_TODELOVERS || nextToken.Type == ADD {
			root.Arguments = append(root.Arguments, a.parser(nextToken))
		} else {
			root.Arguments = append(root.Arguments, a.nodeFactory.Create(nextToken.Type, nextToken.Value))
		}

	}

	return root
}

func (a *ASTAssembler) Assembly(debug bool) *ASTAssembler {
	root := a.parser(a.lexer.NextToken())
	if debug {
		a.debug(root, "")
	}

	a.root = root

	return a

}

func (a *ASTAssembler) GetRoot() Node { return a.root }

func (a *ASTAssembler) debug(node Node, indent string) {
	fmt.Printf("%sType: %s, Token: %v\n", indent, node.Type().String(), node.Token())

	if node.Type() == OPEN_PAREN || node.Type() == CONTEXT || node.Type() == DEF_TODELOVERS || node.Type() == CALL_FUNCTION || node.Type() == ADD {
		for _, child := range node.(*FunctionCallNode).Arguments {
			a.debug(child, indent+"  ")
		}
	}
}
