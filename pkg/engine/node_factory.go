package engine

type nodeFactory struct{}

func (nf *nodeFactory) Create(nodeType tokenType, value interface{}) Node {
	token := value.(string)
	switch nodeType {
	case open_paren:
		return &openParenNode{T: token, Name: value.(string), Nodes: []Node{}}
	case whitespace:
		return &whiteSoaceNode{T: token}
	case context:
		return &contextNode{T: token}
	case newline:
		return &newLineNode{T: token}
	case print:
		return &printNode{T: token}
	case def_todelovers:
		return &defTodeLoverstNode{T: token}
	case types:
		return &typeNode{T: token}
	case leftcol:
		return &leftColNode{T: token}
	case rightcol:
		return &rightColNode{T: token}
	case public:
		return &publicNode{T: token}
	case private:
		return &privateNode{T: token}
	case hashTag:
		return &hashTagNode{T: token}
	case function_zone:
		return &functionZoneNode{T: token}
	case def_func:
		return &functionNode{T: token}
	case main:
		return &mainNode{T: token}
	case leftarroe:
		return &leftArrowNode{T: token}
	case rightarrow:
		return &rightArrowNode{T: token}
	case add:
		return &addNode{T: token}
	case number:
		return &numberNode{T: token}
	case identifier:
		return &identifierNode{T: token}
	case close_paren:
		return &closeParenNode{T: token}
	case func_params:
		return &functionParamNode{T: token}
	case body_func_init:
		return &functionBody{T: token}
	case str:
		return &stringNode{T: token}
	case body_func_eof:
		return &returnTypeNode{T: token}
	case set_variable:
		return &setVariable{T: token}
	case get_variable:
		return &getVariable{T: token}
	case wait:
		return &waitThread{T: token}
	case thread:
		return &newThread{T: token}
	default:
		return &eofNode{T: token}
	}

}

func NewNodeFactory() *nodeFactory { return new(nodeFactory) }
