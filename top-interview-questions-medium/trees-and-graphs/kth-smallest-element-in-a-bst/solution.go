package main

import (
	"fmt"
	. "leetcode/commom"
)

func kthSmallest(root *TreeNode, k int) int {
	var idx int
	return DFS(root, k, &idx)
}

func DFS(root *TreeNode, k int, idx *int) int {
	if root == nil {
		return 0
	}
	res := DFS(root.Left, k, idx)
	if res != 0 {
		return res
	}
	*idx++
	if *idx == k {
		return root.Val
	}
	return DFS(root.Right, k, idx)
}

func kthSmallest2(root *TreeNode, k int) int {
	_, data := loopTree(root, &k)
	return data
}

func loopTree(root *TreeNode, k *int) (bool, int) {
	if root == nil {
		return false, 0
	}

	exist, data := loopTree(root.Left, k)
	if exist {
		return true, data
	}

	(*k)--
	if *k == 0 {
		return true, root.Val
	}

	exist, data = loopTree(root.Right, k)
	if exist {
		return true, data
	}

	return false, 0
}

func main() {
	root := &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 6}}
	r := kthSmallest(root, 3)
	fmt.Println(r)
}
