package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkSubTree(root *TreeNode, min int, max int) bool {
	if root.Val <= min || root.Val >= max {
		return false
	}
	flag := true
	if root.Left != nil {
		flag = checkSubTree(root.Left, min, root.Val)
	}
	if root.Right != nil && flag {
		flag = checkSubTree(root.Right, root.Val, max)
	}
	return flag
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return checkSubTree(root, math.MinInt64, math.MaxInt64)
}
