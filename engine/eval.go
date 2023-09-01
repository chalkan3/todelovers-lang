package engine

import "fmt"

type evaluation struct{}

func NewEvaluation() *evaluation { return new(evaluation) }

func (e *evaluation) Eval(node Node) interface{} {
	result := node.Eval(nil)
	if isNewContext(node.Type()) {
		for _, child := range node.(*functionCallNode).Arguments {
			fmt.Println(child.Eval(nil))
		}
	}

	return result
}
