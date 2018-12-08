package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertNode(nums []int, root *TreeNode) {
	length := len(nums)
	if length != 0 {
		mid := length / 2
		root.Val = nums[mid]

		if length > 1 {
			root.Left = &TreeNode{0, nil, nil}
			insertNode(nums[:mid], root.Left)
		}
		if length > 2 {
			root.Right = &TreeNode{0, nil, nil}
			insertNode(nums[mid+1:], root.Right)
		}
	}
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode{nums[0], nil, nil}
	insertNode(nums, root)
	return root
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	res := sortedArrayToBST(nums)
	fmt.Println(res)
}
