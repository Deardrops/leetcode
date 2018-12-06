package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	start := &ListNode{0, head}
	fast := start
	slow := start
	for ; n >= 0; n-- {
		fast = fast.Next
	}
	for ; fast != nil; fast = fast.Next {
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return start.Next
}

func main() {
	LinkList := ListNode{1, nil}
	res := removeNthFromEnd(&LinkList, 1)
	fmt.Println(res)
}
