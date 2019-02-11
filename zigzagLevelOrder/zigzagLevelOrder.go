package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	level := []*TreeNode{root}
	leftToRight := true
	for len(level) != 0 {
		var subLevel []*TreeNode
		var tmp []int
		if leftToRight {
			for _, node := range level {
				if node == nil {
					continue
				}
				tmp = append(tmp, node.Val)
				subLevel = append(subLevel, node.Left)
				subLevel = append(subLevel, node.Right)
			}
		} else {
			for _, node := range level {
				if node == nil {
					continue
				}
				tmp = append([]int{node.Val}, tmp...)
				subLevel = append(subLevel, node.Left)
				subLevel = append(subLevel, node.Right)
			}
		}
		if len(tmp) != 0 {
			res = append(res, tmp)
		}
		level = subLevel
		leftToRight = !leftToRight
	}
	return res
}
