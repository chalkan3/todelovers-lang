package engine

type Node interface {
	Eval(sTable *SymbolTable) interface{}
	Type() TokenType
	Token() string
}

type NumberNode struct {
	T     string
	Value int
}

func (nn *NumberNode) Type() TokenType { return NUMBER }
func (nn *NumberNode) Token() string   { return nn.T }
func (nn *NumberNode) Eval(sTable *SymbolTable) interface{} {
	return nn.Value
}

type OperationNode struct {
	T        string
	Operator string
	Left     Node
	Right    Node
}

func (on *OperationNode) Type() TokenType { return OPEN_PAREN }
func (on *OperationNode) Token() string   { return on.T }

func (on *OperationNode) Eval(sTable *SymbolTable) interface{} {
	leftValue := on.Left.Eval(sTable).(int)
	rightValue := on.Right.Eval(sTable).(int)

	switch on.Operator {
	case "+":
		return leftValue + rightValue
	case "-":
		return leftValue - rightValue
	case "*":
		return leftValue * rightValue
	case "/":
		if rightValue == 0 {
			panic("Division by zero")
		}
		return leftValue / rightValue
	default:
		panic("Unsupported operator: " + on.Operator)
	}

}

type StringNode struct {
	T     string
	Value string
}

func (sn *StringNode) Type() TokenType { return IDENTIFIER }
func (sn *StringNode) Token() string   { return sn.T }
func (sn *StringNode) Eval(sTable *SymbolTable) interface{} {
	return sn.Value
}

// FunctionCallNode represents a function call in the AST.
type FunctionNode struct {
	T          string
	Name       string
	Parameters []string
	Body       Node
}

func (fn *FunctionNode) Type() TokenType { return DEF_FUNC }
func (fn *FunctionNode) Token() string   { return fn.T }

// FunctionNode represents a function definition in the AST.
func (fn *FunctionNode) Eval(sTable *SymbolTable) interface{} {
	sTable.AddFunction(fn.Name, fn)
	return nil
}

// FunctionCallNode represents a function call in the AST.
type FunctionCallNode struct {
	T         string
	Name      string
	Arguments []Node
}

func (fc *FunctionCallNode) Type() TokenType { return CALL_FUNCTION }
func (fc *FunctionCallNode) Token() string   { return fc.T }
func (fc *FunctionCallNode) Eval(sTable *SymbolTable) interface{} {
	fn := sTable.GetFunction(fc.Name)
	if fn == nil {
		panic("Function not found: " + fc.Name)
	}

	// Create a new symbol table for the function call, inheriting the parent table's functions
	newSymbolTable := NewSymbolTable()
	for name, f := range sTable.Functions {
		newSymbolTable.AddFunction(name, f)
	}

	if len(fc.Arguments) != len(fn.Parameters) {
		panic("Argument count mismatch for function: " + fc.Name)
	}

	// Bind arguments to parameters
	for i, paramName := range fn.Parameters {
		argValue := fc.Arguments[i].Eval(sTable)
		newSymbolTable.Variables[paramName] = argValue

	}

	// Evaluate the function's body within its own symbol table
	return fn.Body.Eval(newSymbolTable)
}

type OpenParenNode struct {
	T string
}

func (op *OpenParenNode) Type() TokenType { return OPEN_PAREN }
func (op *OpenParenNode) Token() string   { return op.T }
func (op *OpenParenNode) Eval(sTable *SymbolTable) interface{} {
	return nil
}

type WhiteSpaceNode struct {
	T string
}

func (w *WhiteSpaceNode) Type() TokenType { return WHITESPACE }
func (w *WhiteSpaceNode) Token() string   { return w.T }
func (w *WhiteSpaceNode) Eval(sTable *SymbolTable) interface{} {
	return nil
}

type ContextNode struct {
	T string
}

func (ct *ContextNode) Type() TokenType { return CONTEXT }
func (ct *ContextNode) Token() string   { return ct.T }
func (ct *ContextNode) Eval(sTable *SymbolTable) interface{} {
	return nil
}

type ReturnNode struct {
	T     string
	Value Node
}

func (r *ReturnNode) Type() TokenType { return RETURN }
func (r *ReturnNode) Token() string   { return r.T }
func (r *ReturnNode) Eval(s *SymbolTable) interface{} {
	return r.Value.Eval(s)
}

type NewLineNode struct {
	T     string
	Value Node
}

func (r *NewLineNode) Type() TokenType { return NEWLINE }
func (r *NewLineNode) Token() string   { return r.T }
func (r *NewLineNode) Eval(s *SymbolTable) interface{} {
	return r.Value.Eval(s)
}

type PrintNode struct {
	T     string
	Value Node
}

func (r *PrintNode) Type() TokenType { return PRINT }
func (r *PrintNode) Token() string   { return r.T }
func (r *PrintNode) Eval(s *SymbolTable) interface{} {
	return r.Value.Eval(s)
}

type DefTodeLoverstNode struct {
	T     string
	Value Node
}

func (r *DefTodeLoverstNode) Type() TokenType { return DEF_TODELOVERS }
func (r *DefTodeLoverstNode) Token() string   { return r.T }
func (r *DefTodeLoverstNode) Eval(s *SymbolTable) interface{} {
	return r.Value.Eval(s)
}

type TypeNode struct {
	T     string
	Value Node
}

func (r *TypeNode) Type() TokenType { return TYPE }
func (r *TypeNode) Token() string   { return r.T }
func (r *TypeNode) Eval(s *SymbolTable) interface{} {
	return r.Value.Eval(s)
}

type LeftColNode struct {
	T     string
	Value Node
}

func (r *LeftColNode) Type() TokenType { return LEFTCOL }
func (r *LeftColNode) Token() string   { return r.T }
func (r *LeftColNode) Eval(s *SymbolTable) interface{} {
	return r.Value.Eval(s)
}

type RightColNode struct {
	T     string
	Value Node
}

func (r *RightColNode) Type() TokenType { return RIGHTCOL }
func (r *RightColNode) Token() string   { return r.T }
func (r *RightColNode) Eval(s *SymbolTable) interface{} {
	return r.Value.Eval(s)
}

type PublicNode struct {
	T     string
	Value Node
}

func (r *PublicNode) Type() TokenType { return PUBLIC }
func (r *PublicNode) Token() string   { return r.T }
func (r *PublicNode) Eval(s *SymbolTable) interface{} {
	return r.Value.Eval(s)
}

type PrivateNode struct {
	T     string
	Value Node
}

func (r *PrivateNode) Type() TokenType { return PRIVATE }
func (r *PrivateNode) Token() string   { return r.T }
func (r *PrivateNode) Eval(s *SymbolTable) interface{} {
	return r.Value.Eval(s)
}

type HashTagNode struct {
	T     string
	Value Node
}

func (r *HashTagNode) Type() TokenType { return HASHTAG }
func (r *HashTagNode) Token() string   { return r.T }
func (r *HashTagNode) Eval(s *SymbolTable) interface{} {
	return r.Value.Eval(s)
}

type FunctionZoneNode struct {
	T     string
	Value Node
}

func (r *FunctionZoneNode) Type() TokenType { return FUNCTION_ZONE }
func (r *FunctionZoneNode) Token() string   { return r.T }
func (r *FunctionZoneNode) Eval(s *SymbolTable) interface{} {
	return r.Value.Eval(s)
}

type MainNode struct {
	T     string
	Value Node
}

func (r *MainNode) Type() TokenType { return MAIN }
func (r *MainNode) Token() string   { return r.T }
func (r *MainNode) Eval(s *SymbolTable) interface{} {
	return r.Value.Eval(s)
}

type LeftArrowNode struct {
	T     string
	Value Node
}

func (r *LeftArrowNode) Type() TokenType { return LEFTARROW }
func (r *LeftArrowNode) Token() string   { return r.T }
func (r *LeftArrowNode) Eval(s *SymbolTable) interface{} {
	return r.Value.Eval(s)
}

type RightArrowNode struct {
	T     string
	Value Node
}

func (r *RightArrowNode) Type() TokenType { return RIGHTARROW }
func (r *RightArrowNode) Token() string   { return r.T }
func (r *RightArrowNode) Eval(s *SymbolTable) interface{} {
	return r.Value.Eval(s)
}

type AddNode struct {
	T     string
	Value Node
}

func (r *AddNode) Type() TokenType { return ADD }
func (r *AddNode) Token() string   { return r.T }
func (r *AddNode) Eval(s *SymbolTable) interface{} {
	return r.Value.Eval(s)
}

type IdentifierNode struct {
	T     string
	Value Node
}

func (r *IdentifierNode) Type() TokenType { return TYPE }
func (r *IdentifierNode) Token() string   { return r.T }
func (r *IdentifierNode) Eval(s *SymbolTable) interface{} {
	return r.Value.Eval(s)
}

type CloseParenNode struct {
	T     string
	Value Node
}

func (r *CloseParenNode) Type() TokenType { return CLOSE_PAREN }
func (r *CloseParenNode) Token() string   { return r.T }
func (r *CloseParenNode) Eval(s *SymbolTable) interface{} {
	return r.Value.Eval(s)
}

type EOFNode struct {
	T     string
	Value Node
}

func (r *EOFNode) Type() TokenType { return EOF }
func (r *EOFNode) Token() string   { return r.T }
func (r *EOFNode) Eval(s *SymbolTable) interface{} {
	return r.Value.Eval(s)
}
