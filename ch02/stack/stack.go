package main

type Stack struct {
	stack []int
}

func (s *Stack) size() int {
	return len(s.stack)
}

func (s *Stack) isEmpty() bool {
	return len(s.stack) == 0
}

func (s *Stack) push(element int) {
	s.stack = append(s.stack, element)
}

func (s *Stack) pop() int {
	tail := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return tail
}

func (s *Stack) top() int {
	return s.stack[len(s.stack)-1]
}

func main() {
	// s := Stack{stack: []int{}}

}
