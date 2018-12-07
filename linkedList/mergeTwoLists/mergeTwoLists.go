package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dump := &ListNode{0, nil}
	p := dump
	x, y := l1, l2
	for ; x != nil && y != nil; p = p.Next {
		if x.Val > y.Val {
			p.Next = y
			y = y.Next
		} else {
			p.Next = x
			x = x.Next
		}
	}
	for x != nil {
		p.Next = x
		p = p.Next
		x = x.Next
	}
	for y != nil {
		p.Next = y
		p = p.Next
		y = y.Next
	}
	return dump.Next
}
