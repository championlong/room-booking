package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	test := &TreeNode{
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
	result := levelOrderBottom(test)
	fmt.Println(result)
}

func levelOrderBottom(root *TreeNode) [][]int {
	var resultArrays [][]int
	if root == nil {
		return resultArrays
	}
	//利用队列进行层序遍历
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var tempArr []int
		for _, node := range queue {
			queue = queue[1:]
			tempArr = append(tempArr, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		resultArrays = append([][]int{tempArr}, resultArrays...)
	}
	return resultArrays
}
