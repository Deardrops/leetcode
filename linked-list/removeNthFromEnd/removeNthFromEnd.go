package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	解题思路：
	用双指针解决，fast指针先移动n次，然后slow与fast指针同时向后移动，直到fast移动到链尾。
	迭代结束时，slow指针指向队列中倒数第n个元素，删除它即可。
	时间复杂度是O(n)，空间复杂度是O(1)
*/

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
