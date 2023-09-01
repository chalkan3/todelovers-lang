package engine

func isNewContext(nodeType tokenType) bool {
	return nodeType == open_paren ||
		nodeType == context ||
		nodeType == def_todelovers ||
		nodeType == call_function ||
		nodeType == add
}

func iseof(nodeType tokenType) bool             { return nodeType == eof }
func isWhiteSpace(nodeType tokenType) bool      { return nodeType == whitespace }
func isCloseParentesis(nodeType tokenType) bool { return nodeType == close_paren }
