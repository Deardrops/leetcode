package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	解题思路：
	用双指针解题，每次迭代，
	slow指针向后移动一个节点，
	fast指针向后移动两个节点，
	直到链表尾，如果两者相遇（两者指向同一内存地址），
	则说明有环，反之亦然。
	空间复杂度O(1)。
*/

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head
	for slow.Next != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}
