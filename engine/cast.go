package engine

func toFunctionNode(node Node) *functionNode           { return node.(*functionNode) }
func toFunctionParamNode(node Node) *functionParamNode { return node.(*functionParamNode) }
