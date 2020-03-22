package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	s := newStack()
	for _, token := range tokens {
		if isOperator(token) {
			right := s.pop()
			left := s.pop()
			res := doOperator(left, right, token)
			s.push(res)
		} else {
			num, _ := strconv.Atoi(token)
			s.push(num)
		}
	}
	return s.pop()
}

func isOperator(str string) bool {
	switch str {
	case "+", "-", "*", "/":
		return true
	default:
		return false
	}
}

func doOperator(left, right int, operator string) int {
	switch operator {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		return left / right
	default:
		panic("")
	}
}

type Stack struct {
	elems []int
}

func newStack() *Stack {
	return &Stack{
		elems: make([]int, 0),
	}
}

func (s *Stack) push(num int) {
	s.elems = append(s.elems, num)
}

func (s *Stack) pop() int {
	elem := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return elem
}

func main() {
	data := []string{"4", "13", "5", "/", "+"}
	res := evalRPN(data)
	fmt.Println(res)
}