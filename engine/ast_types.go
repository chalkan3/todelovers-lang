package engine

import (
	"fmt"
	"strconv"
	"strings"
)

type Node interface {
	Eval(sTable *symbolTable) interface{}
	Type() tokenType
	Print(input string)
	Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token)
	GenerateIntermediateCode() string
	Token() string
	String(ident string) string
}

type numberNode struct {
	T     string
	Value int
}

func (nn *numberNode) Type() tokenType { return number }
func (nn *numberNode) Print(input string) {
	fmt.Println(input)
}
func (nn *numberNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {
	number, _ := strconv.Atoi(n.Token())
	nn.Value = number
}
func (nn *numberNode) GenerateIntermediateCode() string {
	return fmt.Sprintf("INT %v", nn.Value)
}
func (nn *numberNode) Token() string { return nn.T }
func (nn *numberNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, nn.Type().String(), nn.Token())
}
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
func (on *operationNode) Print(input string) {
	fmt.Println(input)
}
func (on *operationNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {

}
func (on *operationNode) GenerateIntermediateCode() string {
	return ""
}
func (on *operationNode) Token() string { return on.T }
func (on *operationNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, on.Type().String(), on.Token())
}

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
func (sn *stringNode) Print(input string) {
	fmt.Println(input)
}
func (sn *stringNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {
	sn.Value = n.Token()
}
func (sn *stringNode) GenerateIntermediateCode() string {
	return fmt.Sprintf("STRING %v", sn.Value)
}
func (sn *stringNode) Token() string { return sn.T }
func (sn *stringNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, sn.Type().String(), sn.Token())
}
func (sn *stringNode) Eval(sTable *symbolTable) interface{} {
	return sn.Value
}

// functionCallNode represents a function call in the AST.

type functionParamNode struct {
	T     string
	Name  string
	Value string
	Types string
}

func (fn *functionParamNode) Eval(sTable *symbolTable) interface{} { return nil }
func (fn *functionParamNode) Type() tokenType                      { return func_params }
func (fn *functionParamNode) Print(input string) {
	fmt.Println(input)
}
func (fn *functionParamNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {
	// if > 0
	variableType := strings.Split(n.Token(), "::")
	fn.Value = variableType[0]
	fn.Types = variableType[1]

}
func (fn *functionParamNode) GenerateIntermediateCode() string {
	variableType := strings.Split(fn.Token(), "::")
	fn.Value = variableType[0]
	fn.Types = variableType[1]
	return fmt.Sprintf("PARAM %v %v ", fn.Types, fn.Value)
}
func (fn *functionParamNode) Token() string { return fn.T }
func (fn *functionParamNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, fn.Type().String(), fn.Token())
}

type functionNode struct {
	T          string
	Name       string
	Parameters []Node
	Body       []Node
}

func (fn *functionNode) Type() tokenType { return def_func }
func (fn *functionNode) Print(input string) {
	fmt.Println(input)
}
func (fn *functionNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {
	if n.Type() == leftarroe {
		fn.Parameters = append(fn.Parameters, parse(nextToken))
	}
	nextToken = jumpTokenAndGetNewToken(getCurrentToken(), getNextToken)

	if isNewContext(nextToken.Type) {
		fn.Body = append(fn.Body, parse(nextToken))
	} else {
		fn.Body = append(fn.Body, nodeFactory.Create(nextToken.Type, nextToken.Value))
	}

	fn.Name = "test"

}
func (fn *functionNode) GenerateIntermediateCode() string {
	code := `
$FUNC %s(%s)
%s
FUNC$
`
	params := ""
	for _, param := range fn.Parameters {
		params += param.GenerateIntermediateCode()
	}

	bodyCode := ""

	for _, body := range fn.Body {
		bodyCode = body.GenerateIntermediateCode()
	}

	return fmt.Sprintf(code, fn.Name, params, bodyCode)
}
func (fn *functionNode) Token() string { return fn.T }
func (fn *functionNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, fn.Type().String(), fn.Token())
}

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
func (fc *functionCallNode) Print(input string) {
	fmt.Println(input)
	for _, child := range fc.Arguments {
		child.Print(child.String(" "))
	}
}
func (fc *functionCallNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {
	if isNewContext(n.Type()) {
		fc.Arguments = append(fc.Arguments, parse(nextToken))
	} else {
		fc.Arguments = append(fc.Arguments, n)
	}
}
func (fc *functionCallNode) GenerateIntermediateCode() string {

	argumentsCode := ""
	for _, child := range fc.Arguments {
		argumentsCode += child.GenerateIntermediateCode() + " "

	}
	return fmt.Sprintf("call %s(%s)", fc.Name, argumentsCode)
}
func (fc *functionCallNode) Token() string { return fc.T }
func (fc *functionCallNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, fc.Type().String(), fc.Token())
}
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
	// for i, paramName := range fn.Parameters {
	// 	argValue := fc.Arguments[i].Eval(sTable)
	// 	newsymbolTable.Variables[paramName] = argValue

	// }

	// Evaluate the function's body within its own symbol table
	return 0
}

type openParenNode struct {
	T     string
	Name  string
	Nodes []Node
}

func (op *openParenNode) Type() tokenType { return open_paren }
func (op *openParenNode) Print(input string) {
	fmt.Println(input)
	for _, child := range op.Nodes {
		child.Print(child.String(" "))
	}
}
func (op *openParenNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {
	if isNewContext(n.Type()) {
		op.Nodes = append(op.Nodes, parse(nextToken))

	} else {
		op.Nodes = append(op.Nodes, n)

	}

}
func (op *openParenNode) GenerateIntermediateCode() string {
	return ""
}
func (op *openParenNode) Token() string { return op.T }
func (op *openParenNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, op.Type().String(), op.Token())
}
func (op *openParenNode) Eval(sTable *symbolTable) interface{} {
	return nil
}

type whiteSoaceNode struct {
	T string
}

func (w *whiteSoaceNode) Type() tokenType { return whitespace }
func (w *whiteSoaceNode) Print(input string) {
	fmt.Println(input)
}
func (w *whiteSoaceNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {

}
func (w *whiteSoaceNode) GenerateIntermediateCode() string {
	return ""
}
func (w *whiteSoaceNode) Token() string { return w.T }
func (w *whiteSoaceNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, w.Type().String(), w.Token())
}
func (w *whiteSoaceNode) Eval(sTable *symbolTable) interface{} {
	return nil
}

type contextNode struct {
	T     string
	Nodes []Node
}

func (ct *contextNode) Type() tokenType { return context }
func (ct *contextNode) Print(input string) {
	fmt.Println(input)
	for _, child := range ct.Nodes {
		child.Print(child.String(" "))
	}
}
func (ct *contextNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {
	if isNewContext(n.Type()) {
		ct.Nodes = append(ct.Nodes, parse(nextToken))

	} else {
		ct.Nodes = append(ct.Nodes, n)

	}
}
func (ct *contextNode) GenerateIntermediateCode() string {
	code := `$INIT_TODE_BROADCAST
%s
TODE_BROADCAST$
`
	childCode := ""
	for _, child := range ct.Nodes {
		childCode += child.GenerateIntermediateCode()
	}
	return fmt.Sprintf(code, childCode)
}
func (ct *contextNode) Token() string { return ct.T }
func (ct *contextNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, ct.Type().String(), ct.Token())
}
func (ct *contextNode) Eval(sTable *symbolTable) interface{} {
	return nil
}

type returnNode struct {
	T     string
	Value Node
}

func (r *returnNode) Type() tokenType { return returns }
func (r *returnNode) Print(input string) {
	fmt.Println(input)
}
func (r *returnNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {

}
func (r *returnNode) GenerateIntermediateCode() string {
	return ""
}
func (r *returnNode) Token() string { return r.T }
func (r *returnNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *returnNode) Eval(s *symbolTable) interface{} {
	return r.Value.Eval(s)
}

type newLineNode struct {
	T     string
	Value Node
}

func (r *newLineNode) Type() tokenType { return newline }
func (r *newLineNode) Print(input string) {
	fmt.Println(input)
}
func (r *newLineNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {

}
func (r *newLineNode) GenerateIntermediateCode() string {
	return ""
}
func (r *newLineNode) Token() string { return r.T }
func (r *newLineNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *newLineNode) Eval(s *symbolTable) interface{} {
	return r.Value.Eval(s)
}

type printNode struct {
	T     string
	Value Node
}

func (r *printNode) Type() tokenType { return print }
func (r *printNode) Print(input string) {
	fmt.Println(input)
}
func (r *printNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {

}
func (r *printNode) GenerateIntermediateCode() string {
	return ""
}
func (r *printNode) Token() string { return r.T }
func (r *printNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *printNode) Eval(s *symbolTable) interface{} {
	return r.Value.Eval(s)
}

type defTodeLoverstNode struct {
	T     string
	Nodes []Node
}

func (r *defTodeLoverstNode) Type() tokenType { return def_todelovers }
func (r *defTodeLoverstNode) Print(input string) {
	fmt.Println(input)
}
func (r *defTodeLoverstNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {
	if isNewContext(n.Type()) {
		r.Nodes = append(r.Nodes, parse(nextToken))
	} else {
		r.Nodes = append(r.Nodes, n)
	}
}
func (r *defTodeLoverstNode) GenerateIntermediateCode() string {
	return ""
}
func (r *defTodeLoverstNode) Token() string { return r.T }
func (r *defTodeLoverstNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *defTodeLoverstNode) Eval(s *symbolTable) interface{} {
	return 0
}

type typeNode struct {
	T     string
	Value Node
}

func (r *typeNode) Type() tokenType { return types }
func (r *typeNode) Print(input string) {
	fmt.Println(input)
}
func (r *typeNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {

}
func (r *typeNode) GenerateIntermediateCode() string {
	return ""
}
func (r *typeNode) Token() string { return r.T }
func (r *typeNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *typeNode) Eval(s *symbolTable) interface{} {
	return r.Value.Eval(s)
}

type leftColNode struct {
	T     string
	Value Node
}

func (r *leftColNode) Type() tokenType { return leftcol }
func (r *leftColNode) Print(input string) {
	fmt.Println(input)
}
func (r *leftColNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {

}
func (r *leftColNode) GenerateIntermediateCode() string {
	return ""
}
func (r *leftColNode) Token() string { return r.T }
func (r *leftColNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *leftColNode) Eval(s *symbolTable) interface{} {
	return r.Value.Eval(s)
}

type rightColNode struct {
	T     string
	Value Node
}

func (r *rightColNode) Type() tokenType { return rightcol }
func (r *rightColNode) Print(input string) {
	fmt.Println(input)
}
func (r *rightColNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {

}
func (r *rightColNode) GenerateIntermediateCode() string {
	return ""
}
func (r *rightColNode) Token() string { return r.T }
func (r *rightColNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *rightColNode) Eval(s *symbolTable) interface{} {
	return r.Value.Eval(s)
}

type publicNode struct {
	T     string
	Value Node
}

func (r *publicNode) Type() tokenType { return public }
func (r *publicNode) Print(input string) {
	fmt.Println(input)
}
func (r *publicNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {

}
func (r *publicNode) GenerateIntermediateCode() string {
	return ""
}
func (r *publicNode) Token() string { return r.T }
func (r *publicNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *publicNode) Eval(s *symbolTable) interface{} {
	return r.Value.Eval(s)
}

type privateNode struct {
	T     string
	Value Node
}

func (r *privateNode) Type() tokenType { return private }
func (r *privateNode) Print(input string) {
	fmt.Println(input)
}
func (r *privateNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {

}
func (r *privateNode) GenerateIntermediateCode() string {
	return ""
}
func (r *privateNode) Token() string { return r.T }
func (r *privateNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *privateNode) Eval(s *symbolTable) interface{} {
	return r.Value.Eval(s)
}

type hashTagNode struct {
	T     string
	Value Node
}

func (r *hashTagNode) Type() tokenType { return hashTag }
func (r *hashTagNode) Print(input string) {
	fmt.Println(input)
}
func (r *hashTagNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {

}
func (r *hashTagNode) GenerateIntermediateCode() string {
	return ""
}
func (r *hashTagNode) Token() string { return r.T }
func (r *hashTagNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *hashTagNode) Eval(s *symbolTable) interface{} {
	return r.Value.Eval(s)
}

type functionZoneNode struct {
	T     string
	Nodes []Node
}

func (r *functionZoneNode) Type() tokenType { return function_zone }
func (r *functionZoneNode) Print(input string) {
	fmt.Println(input)
}
func (r *functionZoneNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {
	if isNewContext(n.Type()) {
		r.Nodes = append(r.Nodes, parse(nextToken))

	} else {
		r.Nodes = append(r.Nodes, n)
	}
}
func (r *functionZoneNode) GenerateIntermediateCode() string {
	code := `$FUNCTION_ZONE
%s
FUNCTION_ZONE$`
	functionsCode := ""
	for _, functions := range r.Nodes {
		functionsCode += functions.GenerateIntermediateCode()
	}
	return fmt.Sprintf(code, functionsCode)
}
func (r *functionZoneNode) Token() string { return r.T }
func (r *functionZoneNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *functionZoneNode) Eval(s *symbolTable) interface{} {
	return 0
}

type mainNode struct {
	T     string
	Nodes []Node
}

func (r *mainNode) Type() tokenType { return main }
func (r *mainNode) Print(input string) {
	fmt.Println(input)
}
func (r *mainNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {
	if isNewContext(n.Type()) {
		r.Nodes = append(r.Nodes, parse(nextToken))

	} else {
		r.Nodes = append(r.Nodes, n)
	}
}
func (r *mainNode) GenerateIntermediateCode() string {
	code := `$MAIN_FRANK
%s
MAIN_FRANK$`

	childCode := ""
	for _, child := range r.Nodes {
		childCode += child.GenerateIntermediateCode()
	}

	return fmt.Sprintf(code, childCode)

}
func (r *mainNode) Token() string { return r.T }
func (r *mainNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *mainNode) Eval(s *symbolTable) interface{} {
	return 0
}

type leftArrowNode struct {
	T     string
	Nodes []Node
}

func (r *leftArrowNode) Type() tokenType { return leftarroe }
func (r *leftArrowNode) Print(input string) {
	fmt.Println(input)
}
func (r *leftArrowNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {
	if isNewContext(n.Type()) {
		r.Nodes = append(r.Nodes, parse(nextToken))

	} else {
		r.Nodes = append(r.Nodes, n)
	}
}
func (r *leftArrowNode) GenerateIntermediateCode() string {
	code := ""
	for _, c := range r.Nodes {
		code += c.GenerateIntermediateCode() + ""
	}

	return code
}
func (r *leftArrowNode) Token() string { return r.T }
func (r *leftArrowNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *leftArrowNode) Eval(s *symbolTable) interface{} {
	return 0
}

type rightArrowNode struct {
	T     string
	Nodes []Node
}

func (r *rightArrowNode) Type() tokenType { return rightarrow }
func (r *rightArrowNode) Print(input string) {
	fmt.Println(input)
}
func (r *rightArrowNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {
	if isNewContext(n.Type()) {
		r.Nodes = append(r.Nodes, parse(nextToken))

	} else {
		r.Nodes = append(r.Nodes, n)

	}
}
func (r *rightArrowNode) GenerateIntermediateCode() string {
	code := ""
	for _, param := range r.Nodes {
		code += param.GenerateIntermediateCode()
	}
	return code
}
func (r *rightArrowNode) Token() string { return r.T }
func (r *rightArrowNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *rightArrowNode) Eval(s *symbolTable) interface{} {
	return 0
}

type addNode struct {
	T     string
	Left  Node
	Right Node
}

func (r *addNode) Type() tokenType { return add }
func (r *addNode) Print(input string) {
	fmt.Println(input)
	r.Left.Print(r.Left.String(" "))
	r.Right.Print(r.Right.String(" "))

}
func (r *addNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {
	r.Left = parse(nextToken)
	if r.Left != nil {
		r.Right = parse(nextToken)
	}
}
func (r *addNode) GenerateIntermediateCode() string {
	return fmt.Sprintf("call add %s %s", r.Left.GenerateIntermediateCode(), r.Right.GenerateIntermediateCode())
}
func (r *addNode) Token() string { return r.T }
func (r *addNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *addNode) Eval(s *symbolTable) interface{} {
	return nil
}

type identifierNode struct {
	T     string
	Value Node
}

func (r *identifierNode) Type() tokenType { return types }
func (r *identifierNode) Print(input string) {
	fmt.Println(input)
}
func (r *identifierNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {

}
func (r *identifierNode) GenerateIntermediateCode() string {
	return ""
}
func (r *identifierNode) Token() string { return r.T }
func (r *identifierNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *identifierNode) Eval(s *symbolTable) interface{} {
	return r.Value.Eval(s)
}

type closeParenNode struct {
	T     string
	Value Node
}

func (r *closeParenNode) Type() tokenType { return close_paren }
func (r *closeParenNode) Print(input string) {
	fmt.Println(input)
}
func (r *closeParenNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {

}
func (r *closeParenNode) GenerateIntermediateCode() string {
	return ""
}
func (r *closeParenNode) Token() string { return r.T }
func (r *closeParenNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *closeParenNode) Eval(s *symbolTable) interface{} {
	return r.Value.Eval(s)
}

type eofNode struct {
	T     string
	Value Node
}

func (r *eofNode) Type() tokenType { return eof }
func (r *eofNode) Print(input string) {
	fmt.Println(input)
}
func (r *eofNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {

}
func (r *eofNode) GenerateIntermediateCode() string {
	return ""
}
func (r *eofNode) Token() string { return r.T }
func (r *eofNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *eofNode) Eval(s *symbolTable) interface{} {
	return r.Value.Eval(s)
}
