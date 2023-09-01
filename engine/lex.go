package engine

import (
	"regexp"
	"strings"
)

type tokenType int

func (t tokenType) String() string {
	return [...]string{
		"OPEN_PAREN",
		"WHITESPACE",
		"CONTEXT",
		"NEWLINE",
		"PRINT",
		"DEF_TODELOVERS",
		"TYPES",
		"LEFTCOL",
		"RIGHTCOL",
		"PUBLIC",
		"PRIVATE",
		"hashTag",
		"FUNCTION_ZONE",
		"DEF_FUNC",
		"MAIN",
		"LEFTARROE",
		"RIGHTARROW",
		"ADD",
		"NUMBER",
		"IDENTIFIER",
		"CLOSE_PAREN",
		"CALL_FUNCTION",
		"RETURNS",
		"EOF",
	}[t]

}

const (
	open_paren = iota
	whitespace
	context
	newline
	print
	def_todelovers
	types
	leftcol
	rightcol
	public
	private
	hashTag
	function_zone
	def_func
	main
	leftarroe
	rightarrow
	add
	number
	identifier
	close_paren
	call_function
	returns
	eof
)

type token struct {
	Type  tokenType
	Value string
}

// Lexer represents the lexer for the DSL.
type lexer struct {
	input   string
	current int
	tokens  []token
}

// NewLexer creates a new lexer for the DSL.
func NewLexer(input string) *lexer {
	return &lexer{input, 0, []token{}}
}

// Nexttoken returns the next token in the input.
func (l *lexer) NextToken() token {
	if l.current >= len(l.tokens) {
		return token{eof, ""}
	}
	token := l.tokens[l.current]
	l.current++

	return token
}

// tokenize performs lexical analysis of the DSL and stores the tokens.
func (l *lexer) Tokenize() *lexer {
	tokenPatterns := []struct {
		pattern *regexp.Regexp
		token   tokenType
	}{
		{regexp.MustCompile(`\(`), open_paren},
		{regexp.MustCompile(`[ \t]+`), whitespace},
		{regexp.MustCompile(`tode-broadcast`), context},
		{regexp.MustCompile(`\n`), newline},
		{regexp.MustCompile(`nando-talk`), print},
		{regexp.MustCompile(`def-todelovers`), def_todelovers},
		{regexp.MustCompile(`type`), types},
		{regexp.MustCompile(`\[`), leftcol},
		{regexp.MustCompile(`\]`), rightcol},
		{regexp.MustCompile(`public`), public},
		{regexp.MustCompile(`private`), private},
		{regexp.MustCompile(`#`), hashTag},
		{regexp.MustCompile(`functions`), function_zone},
		{regexp.MustCompile(`def-func`), def_func},
		{regexp.MustCompile(`main-frank`), main},
		{regexp.MustCompile(`->`), leftarroe},
		{regexp.MustCompile(`<-`), rightarrow},
		{regexp.MustCompile(`add`), add},
		{regexp.MustCompile(`\b\d+\b`), number},
		// {regexp.MustCompile(`\b[^(\s]+\b`), identifier},
		{regexp.MustCompile(`\)`), close_paren},
	}

	lines := strings.Split(l.input, "\n")
	for _, line := range lines {
		for _, pattern := range tokenPatterns {
			for _, match := range pattern.pattern.FindAllString(line, -1) {
				l.tokens = append(l.tokens, token{pattern.token, match})
			}
		}
	}

	return l
}
