package engine

import "fmt"

type Evaluation struct{}

func NewEvaluation() *Evaluation { return new(Evaluation) }

func (e *Evaluation) Eval(node Node) interface{} {
	result := node.Eval(nil)
	if isNewContext(node.Type()) {
		for _, child := range node.(*FunctionCallNode).Arguments {
			fmt.Println(child.Eval(nil))
		}
	}

	return result
}
