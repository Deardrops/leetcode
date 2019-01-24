package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	odd, even, evenHead, curr := head, head.Next, head.Next, head.Next.Next
	for {
		odd.Next = curr
		odd = curr
		if curr.Next == nil {
			even.Next = nil
			break
		}
		curr = curr.Next
		even.Next = curr
		even = curr
		if curr.Next == nil {
			odd.Next = nil
			break
		}
		curr = curr.Next
	}
	odd.Next = evenHead
	return head
}

func arrayToLinkList(nums []int) *ListNode {
	head := &ListNode{}
	curr := head
	for _, num := range nums {
		curr.Next = &ListNode{num, nil}
		curr = curr.Next
	}
	return head.Next
}

func linkListToArray(head *ListNode) []int {
	var nums []int
	curr := head
	for curr != nil {
		nums = append(nums, curr.Val)
		curr = curr.Next
	}
	return nums
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	head := arrayToLinkList(nums)
	res := oddEvenList(head)
	fmt.Println(linkListToArray(res))
}
