package main

import "fmt"

func isValid(s string) bool {
	stack := []int32{'0'}
	for _, char := range s {
		top := stack[len(stack)-1]
		if (top == '(' && char == ')') || (top == '{' && char == '}') || (top == '[' && char == ']') {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}
	return len(stack) == 1
}

func main() {
	res := isValid("()")
	fmt.Println(res)
}
