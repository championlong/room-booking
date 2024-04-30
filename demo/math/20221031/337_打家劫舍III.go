package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	test := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
	}
	rob(test)
}

// 状态转移方程
// f(0) = max(left[0], left[1]) + max(right[0], right[1]) 不偷当前节点得到的最大价值
// f(1) = val + left[0] + right[0] 偷当前节点得到的最大价值
func rob(root *TreeNode) int {

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var dfs func(root *TreeNode) []int
	dfs = func(root *TreeNode) []int {
		if root == nil {
			return []int{0, 0}
		}

		left := dfs(root.Left)
		right := dfs(root.Right)

		val1 := max(left[0], left[1]) + max(right[0], right[1]) //不偷当前节点
		val2 := root.Val + left[0] + right[0] //偷当前节点
		return []int{val1, val2}
	}

	dp := dfs(root)
	return max(dp[0], dp[1])
}
