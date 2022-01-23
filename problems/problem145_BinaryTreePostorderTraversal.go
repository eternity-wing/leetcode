package problems

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil{
		return result
	}
	result = append(result,  postorderTraversal(root.Left)...)
	result = append(result,  postorderTraversal(root.Right)...)
	result = append(result, root.Val)
	return result
}