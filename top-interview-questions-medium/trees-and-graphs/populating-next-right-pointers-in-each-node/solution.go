package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
	Next  *TreeNode
}

// Copied from https://leetcode.com/problems/populating-next-right-pointers-in-each-node/discuss/37472

func connect(root *TreeNode) {
	if root == nil {
		return
	}
	pre := root
	for pre.Next != nil {
		cur := pre
		for cur != nil {
			cur.Left.Next = cur.Right
			if cur.Next != nil {
				cur.Right.Next = cur.Next.Left
			}
			cur = cur.Next
		}
		pre = pre.Next
	}
}
