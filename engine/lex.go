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
		"IDENTIFIER",
		"PRINT",
		"DEF_TODELOVERS",
		"FUNCTION_ZONE",
		"DEF_FUNC",
		"MAIN",
		"TYPES",
		"LEFTCOL",
		"RIGHTCOL",
		"PUBLIC",
		"PRIVATE",
		"HASHTAG",
		"LEFTARROE",
		"FUNC_PARAMS",
		"RIGHTARROE",
		"ADD",
		"NUMBER",
		"EOL_FUNC_PARAM",
		"CALL_FUNCTION",
		"CLOSE_PAREN",
		"RETURNS",
		"EOF",
	}[t]

}

const (
	open_paren tokenType = iota
	whitespace
	context
	newline
	identifier
	print
	def_todelovers
	function_zone
	def_func
	main
	types
	leftcol
	rightcol
	public
	private
	hashTag
	leftarroe
	func_params
	rightarrow
	add
	number
	eol_func_param
	call_function
	close_paren
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

func (l *lexer) GetCurrentToken() token {
	if l.current >= len(l.tokens) {
		return token{eof, ""}
	}
	return l.tokens[l.current]
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
		{regexp.MustCompile(`\b[^(\s]+\b`), identifier},
		{regexp.MustCompile(`nando-talk`), print},
		{regexp.MustCompile(`def-todelovers`), def_todelovers},
		{regexp.MustCompile(`functions`), function_zone},
		{regexp.MustCompile(`def-func`), def_func},
		{regexp.MustCompile(`main-frank`), main},
		{regexp.MustCompile(`type`), types},
		{regexp.MustCompile(`\[`), leftcol},
		{regexp.MustCompile(`\]`), rightcol},
		{regexp.MustCompile(`public`), public},
		{regexp.MustCompile(`private`), private},
		{regexp.MustCompile(`#`), hashTag},
		{regexp.MustCompile(`->`), leftarroe},
		{regexp.MustCompile(`\w+::\w+`), func_params},
		{regexp.MustCompile(`<-`), rightarrow},
		{regexp.MustCompile(`add`), add},
		{regexp.MustCompile(`\b\d+\b`), number},
		{regexp.MustCompile(`\|`), eol_func_param},
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
