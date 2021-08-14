package treesandgraphs


func zigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int

	if root == nil {
		return result
	}

	var storeTreeInArray func(nodes []*TreeNode, level int)
	storeTreeInArray = func(nodes []*TreeNode, level int) {
		numberOfNodes := len(nodes)
		if numberOfNodes == 0 {
			return
		}
		var nextLevelNodes []*TreeNode

		result = append(result, make([]int, numberOfNodes))
		for i := numberOfNodes - 1; i >= 0; i-- {
			node := nodes[i]
			result[level] = append(result[level], node.Val)
			if level % 2 == 0 {
				if node.Left != nil {
					nextLevelNodes = append(nextLevelNodes, node.Left)
				}
				if node.Right != nil {
					nextLevelNodes = append(nextLevelNodes, node.Right)
				}
			}else{
				if node.Right != nil {
					nextLevelNodes = append(nextLevelNodes, node.Right)
				}
				if node.Left != nil {
					nextLevelNodes = append(nextLevelNodes, node.Left)
				}
			}
		}
		storeTreeInArray(nextLevelNodes, level+1)
	}
	storeTreeInArray([]*TreeNode{root}, 0)
	return result

}

func RunZigzagLevelOrder() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	zigzagLevelOrder(root)
}
