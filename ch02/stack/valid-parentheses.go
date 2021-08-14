package main

import "fmt"

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

var m map[rune]rune

func main() {
	m = make(map[rune]rune)
	m['('] = ')'
	m['['] = ']'
	m['{'] = '}'

	str := ""
	fmt.Scanln(&str)
	fmt.Println(isValidParentheses(str))
}

func isValidParentheses(str string) bool {
	stack := Stack{stack: []rune{}}
	runes := []rune(str)
	len := len(runes)

	for i := 0; i < len; i++ {
		c := runes[i]

		if _, ok := m[c]; ok { // 左括号
			stack.push(c)
		} else { // 右括号
			if stack.isEmpty() {
				return false
			}

			if c != m[stack.pop()] {
				return false
			}
		}
	}

	return stack.isEmpty()
}
