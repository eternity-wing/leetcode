package problems

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil{
		return result
	}
	result = append(result, root.Val)
	result = append(result,  preorderTraversal(root.Left)...)
	result = append(result,  preorderTraversal(root.Right)...)
	return result
}
