package engine

import (
	"encoding/json"
	"fmt"
)

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

func (a *astAssembler) createRoot(token token) Node {
	token = jumpTokenAndGetNewTokenRoot(token, a.lexer.NextToken)

	root := a.nodeFactory.Create(token.Type, token.Value)

	return root
}

func (a *astAssembler) getValidNextToken() token {
	return jumpTokenAndGetNewToken(a.lexer.NextToken(), a.lexer.NextToken)
}
func (a *astAssembler) fillWithChild(root Node) Node {
	for {
		nextToken := a.getValidNextToken()

		if isEOF(nextToken, root) {
			break
		}

		if isNodeReturn(root) {
			fmt.Println(nextToken)
		}

		child := a.nodeFactory.Create(nextToken.Type, nextToken.Value)
		root.Fill(child, a.parser, nextToken, a.lexer.GetCurrentToken, a.nodeFactory, a.lexer.NextToken)

	}

	fmt.Printf("%#v\n", root)

	return root

}

func (a *astAssembler) parser(token token) Node { return a.fillWithChild(a.createRoot(token)) }

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
	// node.Print(node.String(" "))
	pp(node)

}

func pp(input Node) {
	bb, _ := json.Marshal(input)

	fmt.Println(string(bb))
}
