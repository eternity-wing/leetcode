package treesandgraphs

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil || p == nil {
		return nil
	}
	if p.Right != nil {
		return getMinNode(p.Right)
	}
	return getGreaterAncestor(root,p)
}

func getMinNode(root *TreeNode) *TreeNode {
	for ; root.Left != nil; root = root.Left {
	}
	return root
}

func getGreaterAncestor(root *TreeNode, node *TreeNode) *TreeNode {
	if node == nil || root == nil {
		return nil
	}
	parent := root
	for ; ; {
		if parent.Val < node.Val {
			return nil
		}

		left := parent.Left
		if left != nil && left.Val > node.Val {
			parent = left
		} else {
			break
		}

	}
	return parent
}

func RunInorderSuccessor()  {
}