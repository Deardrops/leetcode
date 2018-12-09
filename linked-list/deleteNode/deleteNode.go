package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, n int) *ListNode {
	dump := &ListNode{0, head}
	for p := dump; p.Next != nil; p = p.Next {
		if p.Next.Val == n {
			p.Next = p.Next.Next
			break
		}
	}
	return dump.Next
}

func main() {
	test := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	deleteNode(test, 2)
}
