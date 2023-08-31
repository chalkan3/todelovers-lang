package engine

// import "fmt"

// type Node struct {
// 	Type     Token
// 	Value    string
// 	Children []*Node
// }

// func (node *Node) Evaluate() int {
//     switch node.Type {
//     case IDENTIFIER:
//         return 0
//     case NUMBER:
//         return int(node.Value)
//     case OPERATOR:
//         // Evaluate the left and right children.
//         leftValue := node.Children[0].Evaluate()
//         rightValue := node.Children[1].Evaluate()

//         // Combine the results of the evaluations.
//         switch node.Value {
//         case "+":
//             return leftValue + rightValue
//         case "-":
//             return leftValue - rightValue
//         case "*":
//             return leftValue * rightValue
//         case "/":
//             return leftValue / rightValue
//         }
//     }

//     // The node is not a leaf node or an operator.
//     return 0
// }

// func Parse(tokens []Token) {
// 	tree := &Node{Type: KEYWORD, Value: "god-power"}

// 	for _, token := range tokens {
// 		node := &Node{Type: token, Value: ""}
// 		tree.Children = append(tree.Children, node)
// 	}

// 	fmt.Printf("%#v", tree)
// }

// // import "fmt"

// // type nodeType int

// // const (
// // 	nodeError nodeType = iota
// // 	nodeNumber
// // 	nodeOperator
// // 	nodeLeftParen
// // 	nodeRightParen
// // 	nodeRightArrow
// // 	nodeGodPower
// // 	nodeCreatures
// // 	nodePrivate
// // 	nodePublic
// // 	nodeTode
// // 	nodeDefFunc
// // 	nodeMain
// // 	nodeGracefulEnd
// // 	nodeDestroy
// // 	nodeExit
// // 	nodeReturn
// // 	nodeTypes
// // 	nodeEOL
// // 	nodeEOF
// // )

// // type node struct {
// // 	typ   nodeType
// // 	value string
// // 	left  *node
// // 	right *node
// // }

// // func Parse(tokens []token) (*node, error) {
// // 	var stack []*node
// // 	var output []*node

// // 	// pop := func() (*node, bool) {
// // 	// 	if len(stack) == 0 {
// // 	// 		return nil, false
// // 	// 	}
// // 	// 	n := stack[len(stack)-1]
// // 	// 	stack = stack[:len(stack)-1]
// // 	// 	return n, true
// // 	// }

// // 	addNode := func(types nodeType, token token, stack []*node) []*node {
// // 		stack = append(stack, &node{typ: types, value: token.value})
// // 		return stack
// // 	}

// // 	//[
// // 	// &node{typ: nodeLeftParenteses, value: (},
// // 	// &node{typ: nodeGodpower, value: god-power},
// // 	// &node{typ: nodeLeftParenteses, value: (},
// // 	// &node{typ: creatures, value: creatures},
// // 	// &node{typ: nodeLeftParenteses, value: (},
// // 	// &node{typ: nodeLeftParenteses, value: (},
// // 	// &node{typ: nodePrivate, value: private},
// // 	// &node{typ: nodeInt, value: int},
// // 	// &node{typ: node#, value: #},
// // 	// &node{typ: node|, value: |},
// // 	// &node{typ: nodestring, value: string},
// // 	// &node{typ: node#, value: #},
// // 	// &node{typ: node|, value: |},

// // 	//]
// // 	//
// // 	//

// // 	for _, token := range tokens {

// // 		switch token.typ {
// // 		case TokenNumber:
// // 			output = addNode(nodeNumber, token, output)
// // 		case TokenOperator:
// // 			output = addNode(nodeOperator, token, output)
// // 		case TokenLeftParen:
// // 			stack = addNode(nodeLeftParen, token, stack)
// // 		case TokenRightParen:

// // 			for len(stack) > 0 && stack[len(stack)-1].typ != nodeLeftParen {

// // 				top := stack[len(stack)-1]
// // 				stack = stack[:len(stack)-1]
// // 				output = append(output, top)
// // 			}

// // 			if len(stack) == 0 || stack[len(stack)-1].typ != nodeLeftParen {
// // 				return nil, fmt.Errorf("Mismatched parentheses")
// // 			}
// // 			stack = stack[:len(stack)-1]
// // 		case TokenRightArrow:
// // 			stack = addNode(nodeRightArrow, token, stack)
// // 		case TokenGodPower:
// // 			stack = addNode(nodeGodPower, token, stack)
// // 		case TokenCreatures:
// // 			stack = addNode(nodeCreatures, token, stack)
// // 		case TokenPrivate:
// // 			stack = addNode(nodePrivate, token, stack)
// // 		case TokenPublic:
// // 			stack = addNode(nodePublic, token, stack)
// // 		case TokenTode:
// // 			stack = addNode(nodeTode, token, stack)
// // 		case TokenDefFunc:
// // 			stack = addNode(nodeDefFunc, token, stack)
// // 		case TokenMain:
// // 			stack = addNode(nodeMain, token, stack)
// // 		case TokenGracefulEnd:
// // 			stack = addNode(nodeGracefulEnd, token, stack)
// // 		case TokenDestroy:
// // 			stack = addNode(nodeDestroy, token, stack)
// // 		case TokenExit:
// // 			stack = addNode(nodeExit, token, stack)
// // 		case TokenReturn:
// // 			stack = addNode(nodeReturn, token, stack)
// // 		case TokenTypes:
// // 			stack = addNode(nodeTypes, token, stack)
// // 		case TokenEOL:
// // 			stack = addNode(nodeEOL, token, stack)
// // 		case TokenEOF:
// // 			stack = addNode(nodeEOF, token, stack)
// // 		}

// // 	}

// // 	// for _, v := range output {
// // 	// 	fmt.Printf("stack %#v\n", v)
// // 	// }
// // 	// fmt.Println("output", output)

// // 	// Convert the output to an AST with scope between parentheses.

// // 	return nil, nil
// // }
