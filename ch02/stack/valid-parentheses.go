package main

type Stack struct {
	stack []rune
}

func (s *Stack) size() int {
	return len(s.stack)
}

func (s *Stack) isEmpty() bool {
	return len(s.stack) == 0
}

func (s *Stack) push(element rune) {
	s.stack = append(s.stack, element)
}

func (s *Stack) pop() rune {
	top := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return top
}

func (s *Stack) top() rune {
	return s.stack[len(s.stack)-1]
}

func main() {

}
