package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	nodes := []*TreeNode{root}
	level := 0
	for {
		nextLevelNodes := make([]*TreeNode, 0)
		nums := make([]int, 0)
		index := 0
		for _, node := range nodes {
			nums = append(nums, node.Val)
			index++
			if node.Left != nil {
				nextLevelNodes = append(nextLevelNodes, node.Left)
			}
			if node.Right != nil {
				nextLevelNodes = append(nextLevelNodes, node.Right)
			}
		}
		res = append(res, nums)
		if len(nextLevelNodes) != 0 {
			nodes = nextLevelNodes
			level++
		} else {
			break
		}
	}
	return res
}

func main() {
	left := &TreeNode{9, nil, nil}
	right := &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}
	root := &TreeNode{3, left, right}
	res := levelOrder(root)
	fmt.Println(res)
}
