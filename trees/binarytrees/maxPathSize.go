package binarytrees

func MaxPathSum(root *TreeNode) int {
	if root == nil{
		return -2000
	}
	leftMax := MaxPathSum(root.Left)
	rightMax := MaxPathSum(root.Right)

	max := root.Val
	paths := []int{leftMax, leftMax + root.Val, leftMax + root.Val + rightMax, root.Val + rightMax, rightMax}

	for _, path := range paths{
		if path > max{
			max = path
		}
	}


	return max
}