package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	解题思路：
	用双指针解决，fast指针每次迭代向后移动两位，直到链尾。
	slow指针每次迭代向后移动一位，并交换扫描过的前后两个节点。
	迭代结束时，slow指针指向队列最中间的元素，并且链表前半部分已经被反转，
	然后逐一比较前半部分和后半部分的每个元素是否相等即可。
	时间复杂度是O(n)，空间复杂度是O(1)
*/

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	var prev *ListNode

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}

	first, second := prev, slow
	if fast != nil {
		second = second.Next
	}

	for second != nil {
		if first.Val != second.Val {
			return false
		}
		first, second = first.Next, second.Next
	}
	return true
}
