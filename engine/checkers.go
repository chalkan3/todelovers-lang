package engine

func isNewContext(nodeType tokenType) bool {
	return nodeType == context ||
		nodeType == def_todelovers ||
		nodeType == call_function ||
		nodeType == add ||
		nodeType == function_zone ||
		nodeType == main ||
		nodeType == def_func ||
		nodeType == rightarrow ||
		nodeType == body_func_init
}

func jumpToken(token token) bool {
	return isTokenWhiteSpace(token) ||
		isTokenIndentifier(token) ||
		isTokenOpenParentesis(token) ||
		isTokenNewLine(token) ||
		isTokenEOF(token)
}

//remove it
func jumpTokenRoot(token token) bool {
	return isTokenWhiteSpace(token) ||
		isTokenIndentifier(token) ||
		isTokenOpenParentesis(token) ||
		isTokenNewLine(token) ||
		isTokenEOF(token)
	// isTokenBodyFuncEOL(token)

}

func isTokenOpenParentesis(token token) bool   { return token.Type == open_paren }
func isTokenEOF(token token) bool              { return token.Type == eof }
func isTokenWhiteSpace(token token) bool       { return token.Type == whitespace }
func isTokenCloseParentesis(token token) bool  { return token.Type == close_paren }
func isTokenIndentifier(token token) bool      { return token.Type == identifier }
func isTokenFuncParamEOL(token token) bool     { return token.Type == eol_func_param }
func isTokenBodyFuncEOL(token token) bool      { return token.Type == body_func_eof }
func isTokenNewLine(token token) bool          { return token.Type == newline }
func isTokenFunctionBodyInit(token token) bool { return token.Type == body_func_init }
func isTokenBroadCast(token token) bool        { return token.Type == context }
func isTokenDefTodeLovers(token token) bool    { return token.Type == def_todelovers }
func isTokenNumber(token token) bool           { return token.Type == number }
func isTokenMain(token token) bool             { return token.Type == main }
func isFunctionBlockEnd(token token) bool      { return token.Type == func_definition_end }

func isNodeAdd(n Node) bool           { return n.Type() == add }
func isNodeDefFunc(n Node) bool       { return n.Type() == def_func }
func isNodeReturn(n Node) bool        { return n.Type() == rightarrow }
func isFunctionZone(n Node) bool      { return n.Type() == function_zone }
func isFunctionBodyInit(n Node) bool  { return n.Type() == body_func_init }
func isNodeBroadCast(n Node) bool     { return n.Type() == context }
func isNodeDefToadLovers(n Node) bool { return n.Type() == def_todelovers }
func isNodeFuncParam(n Node) bool     { return n.Type() == leftarroe }
func isNodeNumber(n Node) bool        { return n.Type() == number }
func isNodeMain(n Node) bool          { return n.Type() == main }

func isAddEOF(token token, n Node) bool  { return isTokenCloseParentesis(token) && isNodeAdd(n) }
func isMainEOF(token token, n Node) bool { return isTokenCloseParentesis(token) && isNodeMain(n) }
func isFuncParamEOF(token token, n Node) bool {
	return isTokenFuncParamEOL(token) && isNodeFuncParam(n)
}
func isTodeBroadCastEOF(token token, n Node) bool {
	return isTokenCloseParentesis(token) && isNodeBroadCast(n)
}
func isTodeDefTodeLoversEOF(token token, n Node) bool {
	return isTokenCloseParentesis(token) && isNodeDefToadLovers(n)
}

func isFunctionZoneEOF(token token, n Node) bool {
	return isTokenCloseParentesis(token) && isFunctionZone(n)
}
func isDefFuncEOF(token token, n Node) bool {
	return isTokenCloseParentesis(token) && isNodeDefFunc(n)
}
func isReturnEOF(token token, n Node) bool { return isTokenBodyFuncEOL(token) && isNodeReturn(n) }

func isFunctionBodyInitEOF(token token, n Node) bool {
	return isFunctionBodyInit(n) && isFunctionBlockEnd(token)
}

func isEOF(token token, n Node) bool {
	return isAddEOF(token, n) ||
		isReturnEOF(token, n) ||
		isDefFuncEOF(token, n) ||
		isFunctionZoneEOF(token, n) ||
		isTodeBroadCastEOF(token, n) ||
		isTodeDefTodeLoversEOF(token, n) ||
		isFunctionBodyInitEOF(token, n) ||
		isFuncParamEOF(token, n) ||
		isMainEOF(token, n)
}
func jumpTokenAndGetNewToken(token token, next func() token) token {
	for {
		if !jumpToken(token) {
			break
		}

		token = next()

	}

	return token

}

func jumpTokenAndGetNewTokenRoot(token token, next func() token) token {
	for {
		if !jumpTokenRoot(token) {
			break
		}

		token = next()

	}

	return token

}
