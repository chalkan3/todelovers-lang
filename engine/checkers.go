package engine

func isNewContext(nodeType tokenType) bool {
	return nodeType == context ||
		nodeType == def_todelovers ||
		nodeType == call_function ||
		nodeType == add ||
		nodeType == function_zone ||
		nodeType == main ||
		nodeType == def_func

}

func jumpToken(types tokenType) bool {
	return isWhiteSpace(types) ||
		isOpenParentesis(types) ||
		isIndentifier(types) ||
		isFuncEOL(types) ||
		iseof(types) ||
		types == rightarrow
}

func isOpenParentesis(nodeType tokenType) bool  { return nodeType == open_paren }
func iseof(nodeType tokenType) bool             { return nodeType == eof }
func isWhiteSpace(nodeType tokenType) bool      { return nodeType == whitespace }
func isCloseParentesis(nodeType tokenType) bool { return nodeType == close_paren }
func isIndentifier(nodeType tokenType) bool     { return nodeType == identifier }
func isFuncEOL(nodeType tokenType) bool         { return nodeType == eol_func_param }

func ensureFirstRootToken(token token, next func() token) token {
	if !jumpToken(token.Type) {
		return token
	}

	return ensureFirstRootToken(next(), next)

}

func jumpTokenAndGetNewToken(token token, next func() token) token {
	for {
		if !jumpToken(token.Type) {
			break
		}

		token = next()

	}

	return token

}
