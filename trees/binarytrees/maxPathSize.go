package binarytrees

import "math"

func maxPathSum(root *TreeNode) int {
	res := math.MinInt32
	helper(root, &res)
	return res
}

func helper(node *TreeNode, res *int) int {
	if node == nil { return 0 }
	left := helper(node.Left, res)
	right := helper(node.Right, res)
	currentNodeNotAsRoot := max(max(left, right) + node.Val, node.Val)
	currentNodeAsRoot := max(currentNodeNotAsRoot, left + right + node.Val)
	*res = max(*res, currentNodeAsRoot)
	return currentNodeNotAsRoot
}

func max(a, b int) int {
	if a > b { return a }
	return b
}