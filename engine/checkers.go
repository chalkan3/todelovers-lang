package engine

func isNewContext(nodeType TokenType) bool {
	return nodeType == OPEN_PAREN ||
		nodeType == CONTEXT ||
		nodeType == DEF_TODELOVERS ||
		nodeType == CALL_FUNCTION ||
		nodeType == ADD
}

func isEOF(nodeType TokenType) bool             { return nodeType == EOF }
func isWhiteSpace(nodeType TokenType) bool      { return nodeType == WHITESPACE }
func isCloseParentesis(nodeType TokenType) bool { return nodeType == CLOSE_PAREN }
