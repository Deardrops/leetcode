package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	head := &ListNode{}
	curr := head
	for l1 != nil && l2 != nil {
		curr.Next = &ListNode{}
		curr = curr.Next
		num := l1.Val + l2.Val + carry
		remainder := num % 10
		carry = (num - remainder) / 10
		curr.Val = remainder
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		curr.Next = &ListNode{}
		curr = curr.Next
		num := l1.Val + carry
		remainder := num % 10
		carry = (num - remainder) / 10
		curr.Val = remainder
		l1 = l1.Next
	}
	for l2 != nil {
		curr.Next = &ListNode{}
		curr = curr.Next
		num := l2.Val + carry
		remainder := num % 10
		carry = (num - remainder) / 10
		curr.Val = remainder
		l2 = l2.Next
	}
	if carry > 0 {
		curr.Next = &ListNode{}
		curr = curr.Next
		curr.Val = carry
	}
	return head.Next
}

func main() {
	l1 := &ListNode{5, nil}
	l2 := &ListNode{5, nil}
	res := addTwoNumbers(l1, l2)
	fmt.Println(res)
}
