package stack

type Stack struct {
	a   []interface{}
	size uint
}

func NewStack(size uint) *Stack {
	return &Stack{a: make([]interface{}, size), size:0}
}
func (s *Stack) Push(content interface{}) {
	s.a[s.size] = content
	s.size = s.size + 1
	// s.a[size++] = a
}
func (s *Stack) Pop() interface{} {
	s.size = s.size - 1
	return s.a[s.size]
}
func (s *Stack) Size() int {
	return int(s.size)
}
func (s *Stack) isEmpty() bool {
	return s.size == 0
}
