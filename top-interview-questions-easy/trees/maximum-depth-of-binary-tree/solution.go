package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 16ms, beat 12.46%
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}

// 8ms, beat 100%
func maxDepth2(root *TreeNode) int {
	if root != nil {
		leftDepth := maxDepth2(root.Left)
		rightDepth := maxDepth2(root.Right)
		if leftDepth > rightDepth {
			return 1 + leftDepth
		}
		return 1 + rightDepth
	}
	return 0
}
