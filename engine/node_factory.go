package engine

type NodeFactory struct{}

func (nf *NodeFactory) Create(nodeType TokenType, value interface{}) Node {
	token := value.(string)
	switch nodeType {
	case OPEN_PAREN:
		return &FunctionCallNode{T: token, Name: value.(string), Arguments: []Node{}}
	case WHITESPACE:
		return &WhiteSpaceNode{T: token}
	case CONTEXT:
		return &FunctionCallNode{T: token, Name: value.(string), Arguments: []Node{}}
	case NEWLINE:
		return &NewLineNode{T: token}
	case PRINT:
		return &PrintNode{T: token}
	case DEF_TODELOVERS:
		return &FunctionCallNode{T: token, Name: value.(string), Arguments: []Node{}}
	case TYPE:
		return &TypeNode{T: token}
	case LEFTCOL:
		return &LeftColNode{T: token}
	case RIGHTCOL:
		return &RightColNode{T: token}
	case PUBLIC:
		return &PublicNode{T: token}
	case PRIVATE:
		return &PrivateNode{T: token}
	case HASHTAG:
		return &HashTagNode{T: token}
	case FUNCTION_ZONE:
		return &FunctionZoneNode{T: token}
	case DEF_FUNC:
		return &FunctionNode{T: token}
	case MAIN:
		return &MainNode{T: token}
	case LEFTARROW:
		return &LeftArrowNode{T: token}
	case RIGHTARROW:
		return &RightArrowNode{T: token}
	case ADD:
		return &FunctionCallNode{T: token, Name: value.(string), Arguments: []Node{}}
	case NUMBER:
		return &NumberNode{T: token}
	case IDENTIFIER:
		return &IdentifierNode{T: token}
	case CLOSE_PAREN:
		return &CloseParenNode{T: token}
	default:
		return &EOFNode{T: token}
	}

}

func NewNodeFactory() *NodeFactory { return new(NodeFactory) }
