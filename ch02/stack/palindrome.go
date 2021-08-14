package main

import (
	"fmt"
)

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
	str := ""
	fmt.Scanln(&str)
	fmt.Println(isPalindrome(str))
}

func isPalindrome(str string) bool {
	if str == "" {
		return false
	}

	var next int
	runes := []rune(str)
	len := len(runes)
	fmt.Println(len)
	mid := int(len / 2) // 求字符串的中点

	stack := Stack{stack: []rune{}}
	for i := 0; i < mid; i++ {
		stack.push(runes[i])
	}

	if len%2 == 0 {
		next = mid
	} else {
		next = mid + 1
	}

	for i := next; i < len; i++ {
		if stack.pop() != runes[i] {
			return false
		}
	}

	return true
}
