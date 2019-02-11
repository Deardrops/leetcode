package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := new(TreeNode)
	idx := 0
	buildSubTree(preorder, inorder, &idx, root)
	return root
}

func buildSubTree(preorder []int, inorder []int, idx *int, node *TreeNode) {
	for i, val := range inorder {
		if val == preorder[*idx] {
			node.Val = val
			*idx = *idx + 1
			if i != 0 {
				node.Left = new(TreeNode)
				buildSubTree(preorder, inorder[0:i], idx, node.Left)
			}
			if len(inorder)-i-1 != 0 {
				node.Right = new(TreeNode)
				buildSubTree(preorder, inorder[i+1:], idx, node.Right)
			}
			return
		}
	}
}

func main() {
	preorder := []int{3, 9, 8, 20, 15, 7}
	inorder := []int{8, 9, 3, 15, 20, 7}
	res := buildTree(preorder, inorder)
	fmt.Println(res)
}
