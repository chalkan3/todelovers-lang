package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// Token represents a lexer token.
type Token struct {
	Type  string
	Value string
}

// Lexer represents the lexer for the DSL.
type Lexer struct {
	input   string
	current int
	tokens  []Token
}

// NewLexer creates a new lexer for the DSL.
func NewLexer(input string) *Lexer {
	return &Lexer{input, 0, []Token{}}
}

// NextToken returns the next token in the input.
func (l *Lexer) NextToken() Token {
	if l.current >= len(l.tokens) {
		return Token{"", ""}
	}
	token := l.tokens[l.current]
	l.current++

	return token
}

// Tokenize performs lexical analysis of the DSL and stores the tokens.
func (l *Lexer) Tokenize() {
	tokenPatterns := []struct {
		pattern *regexp.Regexp
		token   string
	}{
		{regexp.MustCompile(`\(`), "OPEN_PAREN"},
		{regexp.MustCompile(`[ \t]+`), "WHITESPACE"},
		{regexp.MustCompile(`tode-broadcast`), "CONTEXT"},
		{regexp.MustCompile(`\n`), "NEWLINE"},
		{regexp.MustCompile(`def-todelovers`), "DEF-TODELOVERS"},
		{regexp.MustCompile(`type`), "TYPE"},
		{regexp.MustCompile(`\[`), "LEFTCOL"},
		{regexp.MustCompile(`\]`), "RIGHTCOL"},
		{regexp.MustCompile(`public`), "PUBLIC"},
		{regexp.MustCompile(`private`), "PRIVATE"},
		{regexp.MustCompile(`#`), "HASHTAG"},
		{regexp.MustCompile(`functions`), "FUNCTIONS"},
		{regexp.MustCompile(`def-func`), "DEF-FUNC"},
		{regexp.MustCompile(`main-frank`), "MAIN"},
		{regexp.MustCompile(`->`), "LEFTARROW"},
		{regexp.MustCompile(`<-`), "RIGHTARROW"},
		{regexp.MustCompile(`add`), "ADD"},
		{regexp.MustCompile(`\b\d+\b`), "NUMBER"},
		{regexp.MustCompile(`\b[^(\s]+\b`), "IDENTIFIER"},
		{regexp.MustCompile(`\)`), "CLOSE_PAREN"},
	}

	lines := strings.Split(l.input, "\n")
	fmt.Println(lines)
	for _, line := range lines {
		for _, pattern := range tokenPatterns {
			for _, match := range pattern.pattern.FindAllString(line, -1) {
				l.tokens = append(l.tokens, Token{pattern.token, match})
			}
		}
	}
}

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
		fmt.Printf("Type: %s, Value: %s\n", child.Type, child.Value)
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

	lexer := NewLexer(dsl)
	lexer.Tokenize()
	// printJSON(lexer.tokens)

	// return

	// Function to build an AST node recursively
	var buildAST func(token Token) ASTNode
	buildAST = func(token Token) ASTNode {
		node := ASTNode{Type: token.Type, Value: token.Value}

		for {
			nextToken := lexer.NextToken()

			if nextToken.Type == "" {
				break
			}
			if nextToken.Type == "WHITESPACE" {
				continue
			}

			if nextToken.Type == "CLOSE_PAREN" {
				break
			}

			if nextToken.Type == "OPEN_PAREN" || nextToken.Type == "ADD" {
				t := buildAST(nextToken)
				node.Children = append(node.Children, t)
			} else {
				node.Children = append(node.Children, ASTNode{Type: nextToken.Type, Value: nextToken.Value})
			}

		}

		return node
	}

	// Start building the AST from the top-level
	topLevelNode := buildAST(lexer.NextToken())
	printAST(topLevelNode, "")

	// Evaluate the DSL
	Evaluate(topLevelNode)

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
