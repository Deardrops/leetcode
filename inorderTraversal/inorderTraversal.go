package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	理解一个关键点：函数的返回值可以改写为参数值，只要传递该参数的指针即可
*/
func inorderTraversal(root *TreeNode) []int {
	var res []int
	TraverNode(root, &res)
	return res
}

func TraverNode(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	TraverNode(node.Left, res)
	*res = append(*res, node.Val)
	TraverNode(node.Right, res)
}
