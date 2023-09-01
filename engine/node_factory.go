package engine

type nodeFactory struct{}

func (nf *nodeFactory) Create(nodeType tokenType, value interface{}) Node {
	token := value.(string)
	switch nodeType {
	case open_paren:
		return &functionCallNode{T: token, Name: value.(string), Arguments: []Node{}}
	case whitespace:
		return &whiteSoaceNode{T: token}
	case context:
		return &functionCallNode{T: token, Name: value.(string), Arguments: []Node{}}
	case newline:
		return &newLineNode{T: token}
	case print:
		return &printNode{T: token}
	case def_todelovers:
		return &functionCallNode{T: token, Name: value.(string), Arguments: []Node{}}
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
		return &functionCallNode{T: token, Name: value.(string), Arguments: []Node{}}
	case number:
		return &numberNode{T: token}
	case identifier:
		return &identifierNode{T: token}
	case close_paren:
		return &closeParenNode{T: token}
	default:
		return &eofNode{T: token}
	}

}

func NewNodeFactory() *nodeFactory { return new(nodeFactory) }
