package tvm

type stack struct {
	value []interface{}
}

func newStack() *stack {
	return &stack{
		value: make([]interface{}, 0),
	}
}

func (s *stack) Push(object interface{}) {
	s.value = append(s.value, object)

}
func (s *stack) Pop() interface{} {
	object := s.value[len(s.value)-1]
	s.value = s.value[:len(s.value)-1]
	return object
}
