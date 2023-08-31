package engine

import (
	"regexp"
	"strings"
)

type TokenType int

func (t TokenType) String() string {
	return [...]string{
		"OPEN_PAREN",
		"WHITESPACE",
		"CONTEXT",
		"NEWLINE",
		"PRINT",
		"DEF_TODELOVERS",
		"TYPE",
		"LEFTCOL",
		"RIGHTCOL",
		"PUBLIC",
		"PRIVATE",
		"HASHTAG",
		"FUNCTION_ZONE",
		"DEF_FUNC",
		"MAIN",
		"LEFTARROW",
		"RIGHTARROW",
		"ADD",
		"NUMBER",
		"IDENTIFIER",
		"CLOSE_PAREN",
		"CALL_FUNCTION",
		"RETURN",
		"EOF",
	}[t]
}

const (
	OPEN_PAREN = iota
	WHITESPACE
	CONTEXT
	NEWLINE
	PRINT
	DEF_TODELOVERS
	TYPE
	LEFTCOL
	RIGHTCOL
	PUBLIC
	PRIVATE
	HASHTAG
	FUNCTION_ZONE
	DEF_FUNC
	MAIN
	LEFTARROW
	RIGHTARROW
	ADD
	NUMBER
	IDENTIFIER
	CLOSE_PAREN
	CALL_FUNCTION
	RETURN
	EOF
)

type Token struct {
	Type  TokenType
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
		return Token{EOF, ""}
	}
	token := l.tokens[l.current]
	l.current++

	return token
}

// Tokenize performs lexical analysis of the DSL and stores the tokens.
func (l *Lexer) Tokenize() *Lexer {
	tokenPatterns := []struct {
		pattern *regexp.Regexp
		token   TokenType
	}{
		{regexp.MustCompile(`\(`), OPEN_PAREN},
		{regexp.MustCompile(`[ \t]+`), WHITESPACE},
		{regexp.MustCompile(`tode-broadcast`), CONTEXT},
		{regexp.MustCompile(`\n`), NEWLINE},
		{regexp.MustCompile(`nando-talk`), PRINT},
		{regexp.MustCompile(`def-todelovers`), DEF_TODELOVERS},
		{regexp.MustCompile(`type`), TYPE},
		{regexp.MustCompile(`\[`), LEFTCOL},
		{regexp.MustCompile(`\]`), RIGHTCOL},
		{regexp.MustCompile(`public`), PUBLIC},
		{regexp.MustCompile(`private`), PRIVATE},
		{regexp.MustCompile(`#`), HASHTAG},
		{regexp.MustCompile(`functions`), FUNCTION_ZONE},
		{regexp.MustCompile(`def-func`), DEF_FUNC},
		{regexp.MustCompile(`main-frank`), MAIN},
		{regexp.MustCompile(`->`), LEFTARROW},
		{regexp.MustCompile(`<-`), RIGHTARROW},
		{regexp.MustCompile(`add`), ADD},
		{regexp.MustCompile(`\b\d+\b`), NUMBER},
		// {regexp.MustCompile(`\b[^(\s]+\b`), IDENTIFIER},
		{regexp.MustCompile(`\)`), CLOSE_PAREN},
	}

	lines := strings.Split(l.input, "\n")
	for _, line := range lines {
		for _, pattern := range tokenPatterns {
			for _, match := range pattern.pattern.FindAllString(line, -1) {
				l.tokens = append(l.tokens, Token{pattern.token, match})
			}
		}
	}

	return l
}
