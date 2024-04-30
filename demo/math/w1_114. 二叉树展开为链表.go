package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	if root.Left != nil {
		temp := root.Right
		root.Right = root.Left
		root.Left = nil
		for root.Right != nil {
			root = root.Right
		}
		root.Right = temp
	}
}
