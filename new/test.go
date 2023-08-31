package news

// ASTNode represents an abstract syntax tree node.
type ASTNode struct {
	// The type of the node.
	Type TokenType
	// The children of the node.
	Children []*ASTNode
}

func (node *ASTNode) AddChild(child *ASTNode) {
	node.Children = append(node.Children, child)
}

// KeywordNode represents a keyword AST node.
type KeywordNode struct {
	ASTNode
	// The keyword text.
	Text string
}

// IdentifierNode represents an identifier AST node.
type IdentifierNode struct {
	ASTNode
	// The identifier text.
	Text string
}

// NumberNode represents a number AST node.
type NumberNode struct {
	ASTNode
	// The number value.
	Value int
}

// StringNode represents a string AST node.
type StringNode struct {
	ASTNode
	// The string value.
	Value string
}

// PunctuationNode represents a punctuation AST node.
type PunctuationNode struct {
	ASTNode
	// The punctuation symbol.
	Symbol string
}

// NewKeywordNode creates a new keyword AST node.
func NewKeywordNode(token TokenType) *KeywordNode {
	return &KeywordNode{
		ASTNode: ASTNode{
			Type: token,
		},
		Text: string(token),
	}
}

// NewIdentifierNode creates a new identifier AST node.
func NewIdentifierNode(token TokenType) *IdentifierNode {
	return &IdentifierNode{
		ASTNode: ASTNode{
			Type: token,
		},
		Text: string(token),
	}
}

// NewNumberNode creates a new number AST node.
func NewNumberNode(token TokenType) *NumberNode {
	return &NumberNode{
		ASTNode: ASTNode{
			Type: token,
		},
		Value: int(token),
	}
}

// NewStringNode creates a new string AST node.
func NewStringNode(token TokenType) *StringNode {
	return &StringNode{
		ASTNode: ASTNode{
			Type: token,
		},
		Value: string(token),
	}
}

// NewPunctuationNode creates a new punctuation AST node.
func NewPunctuationNode(token TokenType) *PunctuationNode {
	return &PunctuationNode{
		ASTNode: ASTNode{
			Type: token,
		},
		Symbol: string(token),
	}
}

// TokenType represents the type of a token.
type TokenType int

// The different types of tokens.
const (
	// The start of a new line.
	TokenNewline TokenType = iota
	// A keyword.
	TokenKeyword
	// An identifier.
	TokenIdentifier
	// A literal number.
	TokenNumber
	// A string literal.
	TokenString
	// A punctuation symbol.
	TokenPunctuation
)

// KeywordSet is a set of keywords.
var KeywordSet = map[string]TokenType{
	"context":   TokenKeyword,
	"functions": TokenKeyword,
	"def-func":  TokenKeyword,
	"tode":      TokenKeyword,
	"add":       TokenKeyword,
	"zoia-ae":   TokenKeyword,
}

// PunctuationSet is a set of punctuation symbols.
var PunctuationSet = map[string]TokenType{
	"(":  TokenPunctuation,
	")":  TokenPunctuation,
	"<-": TokenPunctuation,
	",":  TokenPunctuation,
}

// Lex scans a string and returns a slice of tokens.
func Lex(text string) []TokenType {
	tokens := []TokenType{}
	for i := 0; i < len(text); i++ {
		c := text[i]

		// Check if the character is not a whitespace.
		if c != ' ' && c != '\t' && c != '\n' {

			// Check if the character is a keyword.
			if t, ok := KeywordSet[string(c)]; ok {
				tokens = append(tokens, t)
				continue
			}

			// Check if the character is a punctuation symbol.
			if t, ok := PunctuationSet[string(c)]; ok {
				tokens = append(tokens, t)
				continue
			}

			tokens = append(tokens, TokenIdentifier)
		}
	}

	return tokens
}

func Parse(tokens []TokenType) *ASTNode {
	ast := &ASTNode{}
	for _, token := range tokens {
		switch token {
		case TokenKeyword:
			node.AddChild(cast.(*ASTNode, NewKeywordNode(token)))
			ast.AddChild(NewKeywordNode(token))
		case TokenIdentifier:
			ast.AddChild(NewIdentifierNode(token))
		case TokenNumber:
			ast.AddChild(NewNumberNode(token))
		case TokenString:
			ast.AddChild(NewStringNode(token))
		case TokenPunctuation:
			ast.AddChild(NewPunctuationNode(token))
		}
	}

	return ast
}
