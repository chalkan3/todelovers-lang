package engine

type Node interface {
	Eval(sTable *symbolTable) interface{}
	Type() tokenType
	Token() string
}

type numberNode struct {
	T     string
	Value int
}

func (nn *numberNode) Type() tokenType { return number }
func (nn *numberNode) Token() string   { return nn.T }
func (nn *numberNode) Eval(sTable *symbolTable) interface{} {
	return nn.Value
}

type operationNode struct {
	T        string
	Operator string
	Left     Node
	Right    Node
}

func (on *operationNode) Type() tokenType { return open_paren }
func (on *operationNode) Token() string   { return on.T }

func (on *operationNode) Eval(sTable *symbolTable) interface{} {
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

type stringNode struct {
	T     string
	Value string
}

func (sn *stringNode) Type() tokenType { return identifier }
func (sn *stringNode) Token() string   { return sn.T }
func (sn *stringNode) Eval(sTable *symbolTable) interface{} {
	return sn.Value
}

// functionCallNode represents a function call in the AST.
type functionNode struct {
	T          string
	Name       string
	Parameters []string
	Body       Node
}

func (fn *functionNode) Type() tokenType { return def_func }
func (fn *functionNode) Token() string   { return fn.T }

// functionNode represents a function definition in the AST.
func (fn *functionNode) Eval(sTable *symbolTable) interface{} {
	sTable.AddFunction(fn.Name, fn)
	return nil
}

// functionCallNode represents a function call in the AST.
type functionCallNode struct {
	T         string
	Name      string
	Arguments []Node
}

func (fc *functionCallNode) Type() tokenType { return call_function }
func (fc *functionCallNode) Token() string   { return fc.T }
func (fc *functionCallNode) Eval(sTable *symbolTable) interface{} {
	fn := sTable.GetFunction(fc.Name)
	if fn == nil {
		panic("Function not found: " + fc.Name)
	}

	// Create a new symbol table for the function call, inheriting the parent table's functions
	newsymbolTable := NewSymbolTable()
	for name, f := range sTable.Functions {
		newsymbolTable.AddFunction(name, f)
	}

	if len(fc.Arguments) != len(fn.Parameters) {
		panic("Argument count mismatch for function: " + fc.Name)
	}

	// Bind arguments to parameters
	for i, paramName := range fn.Parameters {
		argValue := fc.Arguments[i].Eval(sTable)
		newsymbolTable.Variables[paramName] = argValue

	}

	// Evaluate the function's body within its own symbol table
	return fn.Body.Eval(newsymbolTable)
}

type openParenNode struct {
	T string
}

func (op *openParenNode) Type() tokenType { return open_paren }
func (op *openParenNode) Token() string   { return op.T }
func (op *openParenNode) Eval(sTable *symbolTable) interface{} {
	return nil
}

type whiteSoaceNode struct {
	T string
}

func (w *whiteSoaceNode) Type() tokenType { return whitespace }
func (w *whiteSoaceNode) Token() string   { return w.T }
func (w *whiteSoaceNode) Eval(sTable *symbolTable) interface{} {
	return nil
}

type contextNode struct {
	T string
}

func (ct *contextNode) Type() tokenType { return context }
func (ct *contextNode) Token() string   { return ct.T }
func (ct *contextNode) Eval(sTable *symbolTable) interface{} {
	return nil
}

type returnNode struct {
	T     string
	Value Node
}

func (r *returnNode) Type() tokenType { return returns }
func (r *returnNode) Token() string   { return r.T }
func (r *returnNode) Eval(s *symbolTable) interface{} {
	return r.Value.Eval(s)
}

type newLineNode struct {
	T     string
	Value Node
}

func (r *newLineNode) Type() tokenType { return newline }
func (r *newLineNode) Token() string   { return r.T }
func (r *newLineNode) Eval(s *symbolTable) interface{} {
	return r.Value.Eval(s)
}

type printNode struct {
	T     string
	Value Node
}

func (r *printNode) Type() tokenType { return print }
func (r *printNode) Token() string   { return r.T }
func (r *printNode) Eval(s *symbolTable) interface{} {
	return r.Value.Eval(s)
}

type defTodeLoverstNode struct {
	T     string
	Value Node
}

func (r *defTodeLoverstNode) Type() tokenType { return def_todelovers }
func (r *defTodeLoverstNode) Token() string   { return r.T }
func (r *defTodeLoverstNode) Eval(s *symbolTable) interface{} {
	return r.Value.Eval(s)
}

type typeNode struct {
	T     string
	Value Node
}

func (r *typeNode) Type() tokenType { return types }
func (r *typeNode) Token() string   { return r.T }
func (r *typeNode) Eval(s *symbolTable) interface{} {
	return r.Value.Eval(s)
}

type leftColNode struct {
	T     string
	Value Node
}

func (r *leftColNode) Type() tokenType { return leftcol }
func (r *leftColNode) Token() string   { return r.T }
func (r *leftColNode) Eval(s *symbolTable) interface{} {
	return r.Value.Eval(s)
}

type rightColNode struct {
	T     string
	Value Node
}

func (r *rightColNode) Type() tokenType { return rightcol }
func (r *rightColNode) Token() string   { return r.T }
func (r *rightColNode) Eval(s *symbolTable) interface{} {
	return r.Value.Eval(s)
}

type publicNode struct {
	T     string
	Value Node
}

func (r *publicNode) Type() tokenType { return public }
func (r *publicNode) Token() string   { return r.T }
func (r *publicNode) Eval(s *symbolTable) interface{} {
	return r.Value.Eval(s)
}

type privateNode struct {
	T     string
	Value Node
}

func (r *privateNode) Type() tokenType { return private }
func (r *privateNode) Token() string   { return r.T }
func (r *privateNode) Eval(s *symbolTable) interface{} {
	return r.Value.Eval(s)
}

type hashTagNode struct {
	T     string
	Value Node
}

func (r *hashTagNode) Type() tokenType { return hashTag }
func (r *hashTagNode) Token() string   { return r.T }
func (r *hashTagNode) Eval(s *symbolTable) interface{} {
	return r.Value.Eval(s)
}

type functionZoneNode struct {
	T     string
	Value Node
}

func (r *functionZoneNode) Type() tokenType { return function_zone }
func (r *functionZoneNode) Token() string   { return r.T }
func (r *functionZoneNode) Eval(s *symbolTable) interface{} {
	return r.Value.Eval(s)
}

type mainNode struct {
	T     string
	Value Node
}

func (r *mainNode) Type() tokenType { return main }
func (r *mainNode) Token() string   { return r.T }
func (r *mainNode) Eval(s *symbolTable) interface{} {
	return r.Value.Eval(s)
}

type leftArrowNode struct {
	T     string
	Value Node
}

func (r *leftArrowNode) Type() tokenType { return leftarroe }
func (r *leftArrowNode) Token() string   { return r.T }
func (r *leftArrowNode) Eval(s *symbolTable) interface{} {
	return r.Value.Eval(s)
}

type rightArrowNode struct {
	T     string
	Value Node
}

func (r *rightArrowNode) Type() tokenType { return rightarrow }
func (r *rightArrowNode) Token() string   { return r.T }
func (r *rightArrowNode) Eval(s *symbolTable) interface{} {
	return r.Value.Eval(s)
}

type addNode struct {
	T     string
	Value Node
}

func (r *addNode) Type() tokenType { return add }
func (r *addNode) Token() string   { return r.T }
func (r *addNode) Eval(s *symbolTable) interface{} {
	return r.Value.Eval(s)
}

type identifierNode struct {
	T     string
	Value Node
}

func (r *identifierNode) Type() tokenType { return types }
func (r *identifierNode) Token() string   { return r.T }
func (r *identifierNode) Eval(s *symbolTable) interface{} {
	return r.Value.Eval(s)
}

type closeParenNode struct {
	T     string
	Value Node
}

func (r *closeParenNode) Type() tokenType { return close_paren }
func (r *closeParenNode) Token() string   { return r.T }
func (r *closeParenNode) Eval(s *symbolTable) interface{} {
	return r.Value.Eval(s)
}

type eofNode struct {
	T     string
	Value Node
}

func (r *eofNode) Type() tokenType { return eof }
func (r *eofNode) Token() string   { return r.T }
func (r *eofNode) Eval(s *symbolTable) interface{} {
	return r.Value.Eval(s)
}
