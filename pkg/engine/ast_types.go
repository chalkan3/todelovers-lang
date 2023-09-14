package engine

import (
	"fmt"
	"mary_guica/pkg/interpreter"
	"math/rand"
	"strconv"
	"strings"
)

type Node interface {
	Eval() interface{}
	Type() tokenType
	Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token)
	RegisterSymbols(symbolTable *symbolTable, currentScope *scope)
	Print(input string)
	GenerateIntermediateCode(st *symbolTable) []byte
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
func (nn *numberNode) RegisterSymbols(symbolTable *symbolTable, currentScope *scope) {

}
func (nn *numberNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {
	number, _ := strconv.Atoi(n.Token())
	nn.Value = number
}
func (nn *numberNode) GenerateIntermediateCode(st *symbolTable) []byte {
	return []byte{byte(nn.Value)}
}
func (nn *numberNode) Token() string { return nn.T }
func (nn *numberNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, nn.Type().String(), nn.Token())
}
func (nn *numberNode) Eval() interface{} {
	return nn.Value
}

type stringNode struct {
	T     string
	Value string
}

func (sn *stringNode) Type() tokenType { return str }
func (sn *stringNode) Print(input string) {
	fmt.Println(input)
}
func (sn *stringNode) RegisterSymbols(symbolTable *symbolTable, currentScope *scope) {

}
func (sn *stringNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {
	sn.Value = strings.ReplaceAll(n.Token(), "\"", "")
}
func (sn *stringNode) GenerateIntermediateCode(st *symbolTable) []byte {

	code := []byte{byte(len(sn.Value))}
	for _, c := range sn.Value {
		code = append(code, byte(c))
	}

	return code
}
func (sn *stringNode) Token() string { return sn.T }
func (sn *stringNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, sn.Type().String(), sn.Token())
}
func (sn *stringNode) Eval() interface{} {
	return sn.Value
}

// functionCallNode represents a function call in the AST.

type functionParamNode struct {
	T     string
	Name  string
	Value string
	Types string
}

func (fn *functionParamNode) Eval() interface{} { return nil }
func (fn *functionParamNode) Type() tokenType   { return func_params }
func (fn *functionParamNode) Print(input string) {
	fmt.Println(input)
}
func (fn *functionParamNode) RegisterSymbols(symbolTable *symbolTable, currentScope *scope) {
	sbl := Symblo{
		Name:  fn.Name,
		Type:  fn.Type().String(),
		Value: fn.Value,
	}

	symbolTable.AddSymbol(fmt.Sprintf("PARAM: %s", fn.Name), sbl)
}
func (fn *functionParamNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {
	// if > 0
	variableType := strings.Split(n.Token(), "::")
	fn.Value = variableType[0]
	fn.Types = variableType[1]

}
func (fn *functionParamNode) GenerateIntermediateCode(st *symbolTable) []byte {
	variableType := strings.Split(fn.Token(), "::")
	fn.Value = variableType[0]
	fn.Types = variableType[1]

	code := []byte{byte(len(fn.Types))}
	for _, c := range fn.Types {
		code = append(code, byte(c))
	}
	return code
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
func (fn *functionNode) RegisterSymbols(symbolTable *symbolTable, currentScope *scope) {
	functionScope := NewScope(fn.Name)
	currentScope.AddChildScope(functionScope)
	currentScope = functionScope
	sbl := Symblo{
		Name:    fn.Name,
		Type:    fn.Type().String(),
		Value:   "",
		Address: 0x01,
	}
	symbolTable.AddSymbol(fmt.Sprintf("FUNCTION: %v", fn.Name), sbl)

	for _, paramChild := range fn.Parameters {
		paramChild.RegisterSymbols(symbolTable, currentScope)
	}

	for _, paramChild := range fn.Body {
		paramChild.RegisterSymbols(symbolTable, currentScope)
	}

}
func (fn *functionNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {
	if n.Type() == leftarroe {
		fn.Parameters = append(fn.Parameters, parse(nextToken))
	}

	if n.Type() == body_func_init {
		fn.Body = append(fn.Body, parse(nextToken))
	}

	randomNumber := rand.Intn(100) // Replace 100 with your desired range

	randomString := strconv.Itoa(randomNumber)
	fn.Name = randomString

}
func (fn *functionNode) GenerateIntermediateCode(st *symbolTable) []byte {
	// 	code := `%s:
	// %s
	// %s
	// `
	// 	params := []byte{}
	// 	for _, param := range fn.Parameters {
	// 		params += param.GenerateIntermediateCode()
	// 	}

	// 	bodyCode := []byte{}

	// 	for i, body := range fn.Body {
	// 		bodyCode += body.GenerateIntermediateCode()
	// 		if i < len(fn.Body)-1 {
	// 			bodyCode += "\n"
	// 		}
	// 	}

	return []byte{}
}
func (fn *functionNode) Token() string { return fn.T }
func (fn *functionNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, fn.Type().String(), fn.Token())
}

// functionNode represents a function definition in the AST.
func (fn *functionNode) Eval() interface{} {
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
	for _, child := range fc.Arguments {
		child.Print(child.String(" "))
	}
}
func (fc *functionCallNode) RegisterSymbols(symbolTable *symbolTable, currentScope *scope) {

}
func (fc *functionCallNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {
	if isNewContext(n.Type()) {
		fc.Arguments = append(fc.Arguments, parse(nextToken))
	} else {
		fc.Arguments = append(fc.Arguments, n)
	}
}
func (fc *functionCallNode) GenerateIntermediateCode(st *symbolTable) []byte {

	// argumentsCode := ""
	// for _, child := range fc.Arguments {
	// 	argumentsCode += child.GenerateIntermediateCode() + " "

	// }
	return []byte{}
}
func (fc *functionCallNode) Token() string { return fc.T }
func (fc *functionCallNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, fc.Type().String(), fc.Token())
}
func (fc *functionCallNode) Eval() interface{} {
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
func (op *openParenNode) RegisterSymbols(symbolTable *symbolTable, currentScope *scope) {

}
func (op *openParenNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {
	if isNewContext(n.Type()) {
		op.Nodes = append(op.Nodes, parse(nextToken))

	} else {
		op.Nodes = append(op.Nodes, n)

	}

}
func (op *openParenNode) GenerateIntermediateCode(st *symbolTable) []byte {
	return []byte{}
}
func (op *openParenNode) Token() string { return op.T }
func (op *openParenNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, op.Type().String(), op.Token())
}
func (op *openParenNode) Eval() interface{} {
	return nil
}

type whiteSoaceNode struct {
	T string
}

func (w *whiteSoaceNode) Type() tokenType { return whitespace }
func (w *whiteSoaceNode) Print(input string) {
	fmt.Println(input)
}
func (w *whiteSoaceNode) RegisterSymbols(symbolTable *symbolTable, currentScope *scope) {

}
func (w *whiteSoaceNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {

}
func (w *whiteSoaceNode) GenerateIntermediateCode(st *symbolTable) []byte {
	return []byte{}
}
func (w *whiteSoaceNode) Token() string { return w.T }
func (w *whiteSoaceNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, w.Type().String(), w.Token())
}
func (w *whiteSoaceNode) Eval() interface{} {
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
func (ct *contextNode) RegisterSymbols(symbolTable *symbolTable, currentScope *scope) {
	global := NewScope("global")
	symbolTable.currentScope = global
	currentScope = global
	sbl := Symblo{
		Name: "broadcast",
		Type: ct.Type().String(),
	}

	symbolTable.AddSymbol("FUNCTION: broadcast", sbl)

	for _, child := range ct.Nodes {
		child.RegisterSymbols(symbolTable, currentScope)
	}
}
func (ct *contextNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {
	if isNewContext(n.Type()) {
		ct.Nodes = append(ct.Nodes, parse(nextToken))

	} else {
		ct.Nodes = append(ct.Nodes, n)

	}
}
func (ct *contextNode) GenerateIntermediateCode(st *symbolTable) []byte {

	c := []byte{
		interpreter.CALL, st.GetSymbolAdress("def-todelovers"),
		interpreter.CALL, st.GetSymbolAdress("functions-zone"),
		interpreter.CALL, st.GetSymbolAdress("main"),
	}

	for _, child := range ct.Nodes {
		c = append(c, child.GenerateIntermediateCode(st)...)
	}

	c = append(c, interpreter.HALT)

	return c
}
func (ct *contextNode) Token() string { return ct.T }
func (ct *contextNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, ct.Type().String(), ct.Token())
}
func (ct *contextNode) Eval() interface{} {
	return nil
}

type returnTypeNode struct {
	T     string
	Value string
}

func (r *returnTypeNode) Type() tokenType { return body_func_eof }
func (r *returnTypeNode) Print(input string) {
	fmt.Println(input)
}
func (r *returnTypeNode) RegisterSymbols(symbolTable *symbolTable, currentScope *scope) {

}
func (r *returnTypeNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {
	// if > 0
	types := strings.Split(n.Token(), "::")
	r.Value = types[1]
}
func (r *returnTypeNode) GenerateIntermediateCode(st *symbolTable) []byte {
	// return r.Value
	return []byte{}
}
func (r *returnTypeNode) Token() string { return r.T }
func (r *returnTypeNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *returnTypeNode) Eval() interface{} {
	return r.Value
}

type newLineNode struct {
	T     string
	Value Node
}

func (r *newLineNode) Type() tokenType { return newline }
func (r *newLineNode) Print(input string) {
	fmt.Println(input)
}
func (r *newLineNode) RegisterSymbols(symbolTable *symbolTable, currentScope *scope) {

}
func (r *newLineNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {

}
func (r *newLineNode) GenerateIntermediateCode(st *symbolTable) []byte {
	return []byte{}
}
func (r *newLineNode) Token() string { return r.T }
func (r *newLineNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *newLineNode) Eval() interface{} {
	return 0
}

type printNode struct {
	T     string
	Value Node
}

func (r *printNode) Type() tokenType { return print }
func (r *printNode) Print(input string) {
	fmt.Println(input)
}
func (r *printNode) RegisterSymbols(symbolTable *symbolTable, currentScope *scope) {

}
func (r *printNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {
	if isNewContext(n.Type()) {
		r.Value = parse(nextToken)
	} else {
		n.Fill(n, parse, nextToken, getCurrentToken, nodeFactory, getNextToken)
		r.Value = n
	}
}
func (r *printNode) GenerateIntermediateCode(st *symbolTable) []byte {
	//PRINT R2
	c := []byte{interpreter.LOAD_STRING, 0x02}

	c = append(c, r.Value.GenerateIntermediateCode(st)...)
	c = append(c, []byte{interpreter.PRINT, 0x02}...)

	return c
}
func (r *printNode) Token() string { return r.T }
func (r *printNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *printNode) Eval() interface{} {
	return 0
}

type defTodeLoverstNode struct {
	T     string
	Nodes []Node
}

func (r *defTodeLoverstNode) Type() tokenType { return def_todelovers }
func (r *defTodeLoverstNode) Print(input string) {
	fmt.Println(input)
}
func (r *defTodeLoverstNode) RegisterSymbols(symbolTable *symbolTable, currentScope *scope) {
	functionScope := NewScope("def-todelovers")
	currentScope.AddChildScope(functionScope)
	currentScope = functionScope
	sbl := Symblo{
		Name: "def-todelovers",
		Type: r.Type().String(),
	}

	symbolTable.AddSymbol(fmt.Sprintf("def-todelovers"), sbl)
	for _, child := range r.Nodes {
		child.RegisterSymbols(symbolTable, currentScope)
	}
}
func (r *defTodeLoverstNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {
	if isNewContext(n.Type()) {
		r.Nodes = append(r.Nodes, parse(nextToken))
	} else {
		r.Nodes = append(r.Nodes, n)
	}
}
func (r *defTodeLoverstNode) GenerateIntermediateCode(st *symbolTable) []byte {
	return []byte{}
}
func (r *defTodeLoverstNode) Token() string { return r.T }
func (r *defTodeLoverstNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *defTodeLoverstNode) Eval() interface{} {
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
func (r *typeNode) RegisterSymbols(symbolTable *symbolTable, currentScope *scope) {

}
func (r *typeNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {

}
func (r *typeNode) GenerateIntermediateCode(st *symbolTable) []byte {
	return []byte{}
}
func (r *typeNode) Token() string { return r.T }
func (r *typeNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *typeNode) Eval() interface{} {
	return 0
}

type leftColNode struct {
	T     string
	Value Node
}

func (r *leftColNode) Type() tokenType { return leftcol }
func (r *leftColNode) Print(input string) {
	fmt.Println(input)
}
func (r *leftColNode) RegisterSymbols(symbolTable *symbolTable, currentScope *scope) {

}
func (r *leftColNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {

}
func (r *leftColNode) GenerateIntermediateCode(st *symbolTable) []byte {
	return []byte{}
}
func (r *leftColNode) Token() string { return r.T }
func (r *leftColNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *leftColNode) Eval() interface{} {
	return 0
}

type rightColNode struct {
	T     string
	Value Node
}

func (r *rightColNode) Type() tokenType { return rightcol }
func (r *rightColNode) Print(input string) {
	fmt.Println(input)
}
func (r *rightColNode) RegisterSymbols(symbolTable *symbolTable, currentScope *scope) {

}
func (r *rightColNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {

}
func (r *rightColNode) GenerateIntermediateCode(st *symbolTable) []byte {
	return []byte{}
}
func (r *rightColNode) Token() string { return r.T }
func (r *rightColNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *rightColNode) Eval() interface{} {
	return 0
}

type publicNode struct {
	T     string
	Value Node
}

func (r *publicNode) Type() tokenType { return public }
func (r *publicNode) Print(input string) {
	fmt.Println(input)
}
func (r *publicNode) RegisterSymbols(symbolTable *symbolTable, currentScope *scope) {

}
func (r *publicNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {

}
func (r *publicNode) GenerateIntermediateCode(st *symbolTable) []byte {
	return []byte{}
}
func (r *publicNode) Token() string { return r.T }
func (r *publicNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *publicNode) Eval() interface{} {
	return 0
}

type privateNode struct {
	T     string
	Value Node
}

func (r *privateNode) Type() tokenType { return private }
func (r *privateNode) Print(input string) {
	fmt.Println(input)
}
func (r *privateNode) RegisterSymbols(symbolTable *symbolTable, currentScope *scope) {

}
func (r *privateNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {

}
func (r *privateNode) GenerateIntermediateCode(st *symbolTable) []byte {
	return []byte{}
}
func (r *privateNode) Token() string { return r.T }
func (r *privateNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *privateNode) Eval() interface{} {
	return 0
}

type hashTagNode struct {
	T     string
	Value Node
}

func (r *hashTagNode) Type() tokenType { return hashTag }
func (r *hashTagNode) Print(input string) {
	fmt.Println(input)
}
func (r *hashTagNode) RegisterSymbols(symbolTable *symbolTable, currentScope *scope) {

}
func (r *hashTagNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {

}
func (r *hashTagNode) GenerateIntermediateCode(st *symbolTable) []byte {
	return []byte{}
}
func (r *hashTagNode) Token() string { return r.T }
func (r *hashTagNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *hashTagNode) Eval() interface{} {
	return 0
}

type functionZoneNode struct {
	T     string
	Nodes []Node
}

func (r *functionZoneNode) Type() tokenType { return function_zone }
func (r *functionZoneNode) Print(input string) {
	fmt.Println(input)
}
func (r *functionZoneNode) RegisterSymbols(symbolTable *symbolTable, currentScope *scope) {
	functionScope := NewScope("functions-zone")
	currentScope.AddChildScope(functionScope)
	currentScope = functionScope
	sbl := Symblo{
		Name: "functions-zone",
		Type: r.Type().String(),
	}
	symbolTable.AddSymbol(fmt.Sprintf("functions-zone"), sbl)
	for _, child := range r.Nodes {
		child.RegisterSymbols(symbolTable, currentScope)
	}
}
func (r *functionZoneNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {
	if isNewContext(n.Type()) {
		r.Nodes = append(r.Nodes, parse(nextToken))

	} else {
		r.Nodes = append(r.Nodes, n)
	}
}
func (r *functionZoneNode) GenerateIntermediateCode(st *symbolTable) []byte {
	// 	code := `
	// %s
	// `
	// 	functionsCode := ""
	// 	for _, functions := range r.Nodes {
	// 		functionsCode += functions.GenerateIntermediateCode()
	// 	}
	c := []byte{}
	for _, child := range r.Nodes {
		c = append(c, child.GenerateIntermediateCode(st)...)
	}
	return c
}
func (r *functionZoneNode) Token() string { return r.T }
func (r *functionZoneNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *functionZoneNode) Eval() interface{} {
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
func (r *mainNode) RegisterSymbols(symbolTable *symbolTable, currentScope *scope) {
	functionScope := NewScope("main")
	currentScope.AddChildScope(functionScope)
	currentScope = functionScope

	sbl := Symblo{
		Name: "main",
		Type: r.Type().String(),
	}

	symbolTable.AddSymbol(fmt.Sprintf("main"), sbl)

	for _, child := range r.Nodes {
		child.RegisterSymbols(symbolTable, currentScope)
	}
}
func (r *mainNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {
	if isNewContext(n.Type()) {
		r.Nodes = append(r.Nodes, parse(nextToken))
	} else {
		r.Nodes = append(r.Nodes, n)
	}
}
func (r *mainNode) GenerateIntermediateCode(st *symbolTable) []byte {
	c := []byte{}
	for _, child := range r.Nodes {
		c = append(c, child.GenerateIntermediateCode(st)...)
	}

	return c

}
func (r *mainNode) Token() string { return r.T }
func (r *mainNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *mainNode) Eval() interface{} {
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
func (r *leftArrowNode) RegisterSymbols(symbolTable *symbolTable, currentScope *scope) {
	for _, child := range r.Nodes {
		child.RegisterSymbols(symbolTable, currentScope)
	}
}
func (r *leftArrowNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {
	if isNewContext(n.Type()) {
		r.Nodes = append(r.Nodes, parse(nextToken))

	} else {
		r.Nodes = append(r.Nodes, n)
	}
}
func (r *leftArrowNode) GenerateIntermediateCode(st *symbolTable) []byte {
	// code := ""
	// for _, c := range r.Nodes {
	// 	code += c.GenerateIntermediateCode() + ""
	// }

	return []byte{}
}
func (r *leftArrowNode) Token() string { return r.T }
func (r *leftArrowNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *leftArrowNode) Eval() interface{} {
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
func (r *rightArrowNode) RegisterSymbols(symbolTable *symbolTable, currentScope *scope) {
	for _, child := range r.Nodes {
		child.RegisterSymbols(symbolTable, currentScope)
	}
}
func (r *rightArrowNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {
	if isNewContext(n.Type()) {
		r.Nodes = append(r.Nodes, parse(nextToken))

	} else {
		r.Nodes = append(r.Nodes, n)

	}
}
func (r *rightArrowNode) GenerateIntermediateCode(st *symbolTable) []byte {
	// code := ""
	// for _, param := range r.Nodes {
	// 	code += param.GenerateIntermediateCode()
	// }

	// code += "\nRETURN 3 - -"
	return []byte{}
}
func (r *rightArrowNode) Token() string { return r.T }
func (r *rightArrowNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *rightArrowNode) Eval() interface{} {
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
func (r *addNode) RegisterSymbols(symbolTable *symbolTable, currentScope *scope) {
	r.Right.RegisterSymbols(symbolTable, currentScope)
	r.Left.RegisterSymbols(symbolTable, currentScope)
}
func (r *addNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {
	n.Fill(n, parse, nextToken, getCurrentToken, nodeFactory, getNextToken)
	if r.Left == nil {
		r.Left = n
	} else {
		r.Right = n
	}
}
func (r *addNode) GenerateIntermediateCode(st *symbolTable) []byte {
	// 	_ = `%s
	// %s
	// ADD 1 2 3`
	// 	return fmt.Sprintf(code, r.Left.GenerateIntermediateCode(), r.Right.GenerateIntermediateCode())
	return []byte{}
}
func (r *addNode) Token() string { return r.T }
func (r *addNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *addNode) Eval() interface{} {
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
func (r *identifierNode) RegisterSymbols(symbolTable *symbolTable, currentScope *scope) {
	r.Value.RegisterSymbols(symbolTable, currentScope)
}
func (r *identifierNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {

}
func (r *identifierNode) GenerateIntermediateCode(st *symbolTable) []byte {
	return []byte{}
}
func (r *identifierNode) Token() string { return r.T }
func (r *identifierNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *identifierNode) Eval() interface{} {
	return 0
}

type closeParenNode struct {
	T     string
	Value Node
}

func (r *closeParenNode) Type() tokenType { return close_paren }
func (r *closeParenNode) Print(input string) {
	fmt.Println(input)
}
func (r *closeParenNode) RegisterSymbols(symbolTable *symbolTable, currentScope *scope) {

}
func (r *closeParenNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {

}
func (r *closeParenNode) GenerateIntermediateCode(st *symbolTable) []byte {
	return []byte{}
}
func (r *closeParenNode) Token() string { return r.T }
func (r *closeParenNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *closeParenNode) Eval() interface{} {
	return 0
}

type eofNode struct {
	T     string
	Value Node
}

func (r *eofNode) Type() tokenType { return eof }
func (r *eofNode) Print(input string) {
	fmt.Println(input)
}
func (r *eofNode) RegisterSymbols(symbolTable *symbolTable, currentScope *scope) {

}
func (r *eofNode) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {

}
func (r *eofNode) GenerateIntermediateCode(st *symbolTable) []byte {
	return []byte{}
}
func (r *eofNode) Token() string { return r.T }
func (r *eofNode) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *eofNode) Eval() interface{} {
	return 0

}

type functionBody struct {
	T     string
	Nodes []Node
}

func (r *functionBody) Type() tokenType { return body_func_init }
func (r *functionBody) Print(input string) {
	fmt.Println(input)
}
func (r *functionBody) RegisterSymbols(symbolTable *symbolTable, currentScope *scope) {
	for _, child := range r.Nodes {
		child.RegisterSymbols(symbolTable, currentScope)
	}
}
func (r *functionBody) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {
	if isNewContext(n.Type()) {
		r.Nodes = append(r.Nodes, parse(nextToken))

	} else {
		r.Nodes = append(r.Nodes, n)
	}
}
func (r *functionBody) GenerateIntermediateCode(st *symbolTable) []byte {
	// code := ""
	// for i, f := range r.Nodes {
	// 	code += f.GenerateIntermediateCode()
	// 	if i < len(r.Nodes)-1 {
	// 		code += "\n"
	// 	}
	// }
	return []byte{}
}
func (r *functionBody) Token() string { return r.T }
func (r *functionBody) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *functionBody) Eval() interface{} {
	return 0
}

type setVariable struct {
	T     string
	Right Node
	Left  Node
}

func (r *setVariable) Type() tokenType { return set_variable }
func (r *setVariable) Print(input string) {
	fmt.Println(input)
}
func (r *setVariable) RegisterSymbols(symbolTable *symbolTable, currentScope *scope) {
	sbl := Symblo{
		Name:  r.Left.Token(),
		Type:  r.Type().String(),
		Value: r.Right.Token(),
	}

	symbolTable.AddSymbol(fmt.Sprintf("VARIABLE: %s", r.Left.Token()), sbl)

}
func (r *setVariable) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {
	n.Fill(n, parse, nextToken, getCurrentToken, nodeFactory, getNextToken)
	if r.Left == nil {
		r.Left = n
	} else {
		r.Right = n
	}
}
func (r *setVariable) GenerateIntermediateCode(st *symbolTable) []byte {

	code := []byte{0x0C}
	code = append(code, r.Left.GenerateIntermediateCode(st)...)
	code = append(code, 0x07)
	code = append(code, r.Right.GenerateIntermediateCode(st)...)
	code = append(code, []byte{0x00, 0x08, 0x00, 0xC8}...)

	// return fmt.Sprintf("%s = %s \n", r.Left.GenerateIntermediateCode(), r.Right.GenerateIntermediateCode())
	return code
}
func (r *setVariable) Token() string { return r.T }
func (r *setVariable) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *setVariable) Eval() interface{} {
	return 0
}

type getVariable struct {
	T    string
	Node Node
}

func (r *getVariable) Type() tokenType { return get_variable }
func (r *getVariable) Print(input string) {
	fmt.Println(input)
}
func (r *getVariable) RegisterSymbols(symbolTable *symbolTable, currentScope *scope) {

}
func (r *getVariable) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {
	n.Fill(n, parse, nextToken, getCurrentToken, nodeFactory, getNextToken)
	r.Node = n

}
func (r *getVariable) GenerateIntermediateCode(st *symbolTable) []byte {
	code := []byte{0x0D, 0x02}
	code = append(code, r.Node.GenerateIntermediateCode(st)...)
	// return fmt.Sprintf("%s = %s \n", r.Left.GenerateIntermediateCode(), r.Right.GenerateIntermediateCode())
	return code
}
func (r *getVariable) Token() string { return r.T }
func (r *getVariable) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *getVariable) Eval() interface{} {
	return 0
}

type waitThread struct {
	T     string
	Nodes []Node
}

func (r *waitThread) Type() tokenType { return get_variable }
func (r *waitThread) Print(input string) {
	fmt.Println(input)
}
func (r *waitThread) RegisterSymbols(symbolTable *symbolTable, currentScope *scope) {

}
func (r *waitThread) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {
	if isNewContext(n.Type()) {
		r.Nodes = append(r.Nodes, parse(nextToken))

	} else {
		r.Nodes = append(r.Nodes, n)
	}

}
func (r *waitThread) GenerateIntermediateCode(st *symbolTable) []byte {
	code := []byte{}
	for _, child := range r.Nodes {
		code = append(code, child.GenerateIntermediateCode(st)...)
	}

	code = append(code, interpreter.W_THREAD)
	return code
}
func (r *waitThread) Token() string { return r.T }
func (r *waitThread) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *waitThread) Eval() interface{} {
	return 0
}

type newThread struct {
	T     string
	Nodes []Node
}

func (r *newThread) Type() tokenType { return get_variable }
func (r *newThread) Print(input string) {
	fmt.Println(input)
}
func (r *newThread) RegisterSymbols(symbolTable *symbolTable, currentScope *scope) {

}
func (r *newThread) Fill(n Node, parse func(token token) Node, nextToken token, getCurrentToken func() token, nodeFactory *nodeFactory, getNextToken func() token) {
	if isNewContext(n.Type()) {
		r.Nodes = append(r.Nodes, parse(nextToken))

	} else {
		r.Nodes = append(r.Nodes, n)
	}

}
func (r *newThread) GenerateIntermediateCode(st *symbolTable) []byte {
	code := []byte{interpreter.S_THREAD}

	childrenCode := []byte{}
	for _, child := range r.Nodes {
		childrenCode = append(childrenCode, child.GenerateIntermediateCode(st)...)
	}

	commandSize := byte(len(childrenCode))
	code = append(code, commandSize)
	code = append(code, childrenCode...)
	code = append(code, interpreter.ST_THREAD)
	return code
}
func (r *newThread) Token() string { return r.T }
func (r *newThread) String(ident string) string {
	return fmt.Sprintf("%sType: %s, Token: %v, Name: %s\n", ""+ident, r.Type().String(), r.Token())
}
func (r *newThread) Eval() interface{} {
	return 0
}
