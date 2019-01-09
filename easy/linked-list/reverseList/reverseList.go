package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var next *ListNode
	last := head
	p := head.Next
	head.Next = nil
	for p != nil {
		next = p.Next
		p.Next = last
		last = p
		p = next
	}
	return last
}
func recursiveReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev := head
	p := head.Next
	head.Next = nil
	return swap(prev, p)
}

func swap(prev *ListNode, p *ListNode) *ListNode {
	next := p.Next
	p.Next = prev
	if next != nil {
		return swap(p, next)
	} else {
		return p
	}
}
