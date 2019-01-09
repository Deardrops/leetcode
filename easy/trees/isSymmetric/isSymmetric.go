package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkNode(leftNode *TreeNode, rightNode *TreeNode) bool {
	if leftNode == nil && rightNode == nil {
		return true
	}
	if leftNode != nil && rightNode != nil && leftNode.Val == rightNode.Val {
		return checkNode(leftNode.Left, rightNode.Right) && checkNode(leftNode.Right, rightNode.Left)
	}
	return false
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return checkNode(root.Left, root.Right)
}

func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	stack := make([]*TreeNode, 0)
	leftNode := root.Left
	rightNode := root.Right
	for {
		if leftNode == nil && rightNode == nil {
			if len(stack) == 0 {
				return true
			}
			leftNode = stack[len(stack)-2].Right
			rightNode = stack[len(stack)-1].Left
			stack = stack[:len(stack)-2]
			continue
		}
		if leftNode != nil && rightNode != nil && leftNode.Val == rightNode.Val {
			stack = append(stack, leftNode, rightNode)
			leftNode = leftNode.Left
			rightNode = rightNode.Right
			continue
		}
		return false
	}
}
