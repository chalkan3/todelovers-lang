package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mary_guica/engine"
	"strconv"
)

// ASTNode represents a node in the Abstract Syntax Tree.
type ASTNode struct {
	Type     string
	Value    string
	Children []ASTNode
}

// Eval evaluates the DSL represented by the AST.
func Eval(node ASTNode) int {
	FunctionMap := make(map[string]ASTNode)
	switch node.Type {
	case "NUMBER":
		return parseNumber(node.Value)
	case "ADD":
		// Handle the 'add' operation
		if len(node.Children) < 2 {
			panic("Invalid 'add' operation")
		}
		left := Eval(node.Children[0])
		right := Eval(node.Children[1])
		return left + right
	case "PRINT":
		if len(node.Children) < 2 {
			panic("Invalid 'add' operation")
		}

		left := Eval(node.Children[1])
		right := Eval(node.Children[2])
		fmt.Println("Hello Word", left, right)
		return 0
	case "DEF-FUNC":
		return 0
		// Handle function definitions
		if len(node.Children) != 2 {
			panic("Invalid 'def-func' operation")
		}
		functionName := node.Children[0].Value
		functionBody := node.Children[1]

		// Store the function definition in the FunctionMap
		FunctionMap[functionName] = functionBody

		return 0 // Function definitions have no immediate result
	case "IDENTIFIER":
		// Handle variable references (for simplicity, we assume they are integers)
		return 1
	case "CONTEXT", "FUNCTIONS", "MAIN", "ZOIA-AE":
		// These operations have no immediate effect, so we can simply skip them
		return 0
	default:
		// Handle other node types (e.g., function calls)
		return 0
	}
}

func Evaluate(node ASTNode) {
	for _, child := range node.Children {
		fmt.Printf("Type: %s, Value: %s,\n", child.Type, child.Value)
		fmt.Println(Eval(child))
		if len(child.Children) > 0 {
			Evaluate(child)
		}
	}
}

func parseNumber(value string) int {
	// Parse a string as an integer
	result, err := strconv.Atoi(value)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse integer: %v", err))
	}
	return result
}

func main() {
	dsl, _ := openFile("main.todelovers")
	// dsl := `(context)`

	lexer := engine.NewLexer(dsl).Tokenize()
	nodeFactory := engine.NewNodeFactory()
	assembler := engine.NewASTAssembler(lexer, nodeFactory).Assembly(true)
	_ = assembler
	// engine.Eval(assembler.GetRoot())

	// printJSON(lexer.tokens)

	// return

	// Function to build an AST node recursively
	// var buildAST func(token Token) ASTNode
	// buildAST = func(token Token) ASTNode {
	// 	node := ASTNode{Type: token.Type, Value: token.Value}

	// 	for {
	// 		nextToken := lexer.NextToken()

	// 		if nextToken.Type == "" {
	// 			break
	// 		}
	// 		if nextToken.Type == "WHITESPACE" {
	// 			continue
	// 		}

	// 		if nextToken.Type == "CLOSE_PAREN" {
	// 			break
	// 		}

	// 		if nextToken.Type == "OPEN_PAREN" || nextToken.Type == "ADD" || nextToken.Type == "PRINT" {
	// 			t := buildAST(nextToken)
	// 			node.Children = append(node.Children, t)
	// 		} else {
	// 			node.Children = append(node.Children, ASTNode{Type: nextToken.Type, Value: nextToken.Value})
	// 		}

	// 	}

	// 	return node
	// }

	// // Start building the AST from the top-level
	// topLevelNode := buildAST(lexer.NextToken())
	// printAST(topLevelNode, "")

	// // Evaluate the DSL
	// Evaluate(topLevelNode)

}

func printAST(node ASTNode, indent string) {
	fmt.Printf("%sType: %s, Value: %s\n", indent, node.Type, node.Value)
	for _, child := range node.Children {
		printAST(child, indent+"  ")
	}
}

func printJSON(input interface{}) {
	bb, _ := json.Marshal(input)
	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, bb, "", "\t")
	fmt.Println(string(prettyJSON.Bytes()))
}

func openFile(filename string) (string, error) {
	// Read the file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	// Convert the byte slice to a string
	content := string(data)
	return content, nil
}
