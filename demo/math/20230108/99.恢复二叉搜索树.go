package main

func recoverTree(root *TreeNode) {
	var firstNode, lastNode, latestNode *TreeNode
	var preIter func(node *TreeNode)
	preIter = func(node *TreeNode) {
		if node == nil {
			return
		}
		preIter(node.Left)
		if latestNode != nil && node.Val < latestNode.Val {
			if firstNode == nil {
				firstNode = latestNode
			}
			lastNode = node
		}
		latestNode = node
		preIter(node.Right)
	}
	preIter(root)
	firstNode.Val, lastNode.Val = lastNode.Val, firstNode.Val
}
