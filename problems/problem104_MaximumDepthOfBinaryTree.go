package problems

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil{
		return 0
	}
	return 1 + maxInt(maxDepth(root.Left), maxDepth(root.Right))
}