package engine

import (
	"fmt"
	"strings"
)

// The lexer converts the DSL code into a stream of tokens.
type Lexer struct {
	reader *strings.Reader
}

func NewLexer(reader *strings.Reader) *Lexer {
	return &Lexer{
		reader: reader,
	}
}

func (l *Lexer) NextToken() (token Token, err error) {
	// Read the next character from the input.
	c, err := l.reader.ReadByte()
	if err != nil {
		return EOF, err
	}

	// Switch on the character to determine the token type.
	switch c {
	case '(':
		return LPAREN, nil
	case ')':
		return RPAREN, nil
	case '+':
		return PLUS, nil
	case '-':
		return MINUS, nil
	case '*':
		return MUL, nil
	case '/':
		return DIV, nil
	case ' ':
		return WS, nil
	default:
		if c >= 'a' && c <= 'z' {
			return IDENT, nil
		}
		return INVALID, fmt.Errorf("invalid character: %c", c)
	}
}

// The token type.
type Token int

const (
	EOF Token = iota
	LPAREN
	RPAREN
	PLUS
	MINUS
	MUL
	DIV
	IDENT
	WS
	DEF_FUNC
	CALL_FUNC
	INVALID
)

// The AST represents the abstract syntax tree of the DSL code.
type AST struct {
	Type  Token
	Left  *AST
	Right *AST
}

// The parser converts the stream of tokens into an AST.
type Parser struct {
	lexer *Lexer
}

func NewParser(lexer *Lexer) *Parser {
	return &Parser{
		lexer: lexer,
	}
}

func (a *AST) Evaluate() (int, error) {
    // Recursively evaluate the AST.
    switch a.Type {
    case DEF_FUNC:
        // The function definition is not evaluated.
        return 0, nil
    case CALL_FUNC:
        // The function call is evaluated by calling the function.
        return callFunction(a.Left., a)
    default:
        return 0, fmt.Errorf("unsupported AST type: %v", a.Type)
    }
}

func callFunction(name string, args []*AST) (int, error) {
    // Find the function definition.
    function, err := getFuncDef(name)
    if err != nil {
        return 0, err
    }

    // Evaluate the function arguments.
    argValues := make([]int, len(args))
    for i, arg := range args {
        val, err := arg.Evaluate()
        if err != nil {
            return 0, err
        }
        argValues[i] = val
    }

    // Call the function.
    return function.Body.Evaluate()
}

func getFuncDef(name string) (*AST, error) {
    // Iterate over all the function definitions.
    for _, funcDef := range funcDefs {
        if funcDef.Name == name {
            return funcDef, nil
        }
    }

    // The function definition was not found.
    return nil, fmt.Errorf("function %s not found", name)
}


func (p *Parser) Parse() (*AST, error) {
    // Get the next token.
    token, err := p.lexer.NextToken()
    if err != nil {
        return nil, err
    }

    // Switch on the token type to determine the next step.
    switch token {
    case LPAREN:
        // The next token must be an identifier.
        nextToken, err := p.lexer.NextToken()
        if err != nil {
            return nil, err
        }
        if nextToken != IDENT {
            return nil, fmt.Errorf("expected identifier, got: %s", nextToken)
        }

        // Create a new AST node for the function definition.
        node := &AST{
            Type: DEF_FUNC,
            Name: nextToken.Literal,
        }

        // Parse the function body.
        node.Body, err = p.Parse()
        if err != nil {
            return nil, err
        }

        return node, nil
    case IDENT:
        // The next token must be a function call.
        nextToken, err := p.lexer.NextToken()
        if err != nil {
            return nil, err
        }
        if nextToken != LPAREN {
            return nil, fmt.Errorf("expected function call, got: %s", nextToken)
        }

        // Create a new AST node for the function call.
        node := &AST{
            Type: CALL_FUNC,
            Name: nextToken,
        }

        // Parse the function arguments.
        node.Args, err = p.Parse()
        if err != nil {
            return nil, err
        }

        return node, nil
    default:
        return nil, fmt.Errorf("unexpected token: %s", token)
    }
}

func main() {
    // Create a lexer for the DSL code.
    lexer := NewLexer(strings.NewReader(code))

    // Create a parser for the DSL code.
    parser := NewParser(lexer)

    // Parse the DSL code.
    ast, err := parser.Parse()
    if err != nil {
        fmt.Println(err)
        return
    }

    // Evaluate the DSL code.
    value, err := ast.Evaluate()
    if err != nil {
        fmt.Println(err)
        return
    }

    // Print the value of the DSL code.
    fmt.Println(value)
}

