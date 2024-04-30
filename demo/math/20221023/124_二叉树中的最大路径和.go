package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	test := &TreeNode{
		Val:   -3,
		Left:  nil,
		Right: nil,
	}
	test2 := &TreeNode{
		Val: 9,
		Left: &TreeNode{
			Val:   6,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 13,
			Left: &TreeNode{
				Val:   12,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
		},
	}
	maxPathSum(test)
	maxPathSum(test2)
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32 //取最小的int数，兼容根结点为负数的情况

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := max(0, dfs(root.Left))
		right := max(0, dfs(root.Right))

		temp := root.Val + left + right

		maxSum = max(temp, maxSum)
		return root.Val + max(0, max(left, right))
	}

	dfs(root)
	return maxSum
}
