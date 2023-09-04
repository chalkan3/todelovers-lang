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
func (a *astAssembler) parser(token token) Node {
	token = ensureFirstRootToken(token, a.lexer.NextToken)
	root := a.nodeFactory.Create(token.Type, token.Value)

	for {
		nextToken := a.lexer.NextToken()
		if isWhiteSpace(nextToken.Type) || nextToken.Type == open_paren || nextToken.Type == identifier {
			continue
		}

		if iseof(nextToken.Type) || isCloseParentesis(nextToken.Type) || nextToken.Type == eol_func_param {
			break
		}

		child := a.nodeFactory.Create(nextToken.Type, nextToken.Value)
		root.Fill(child, a.parser, nextToken, a.lexer.GetCurrentToken, a.nodeFactory, a.lexer.NextToken)

	}

	return root
}

func (a *astAssembler) appendFunctionCall(nextToken token, value interface{}, root Node) Node {
	castRoot := root.(*functionCallNode)
	if isNewContext(nextToken.Type) {
		castRoot.Arguments = append(castRoot.Arguments, a.parser(nextToken))
	} else {
		castRoot.Arguments = append(castRoot.Arguments, a.nodeFactory.Create(nextToken.Type, nextToken.Value))
	}

	return castRoot

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
	// node.Print(node.String(" "))
	pp(node)

}

func pp(input Node) {
	bb, _ := json.Marshal(input)

	fmt.Println(string(bb))
}
