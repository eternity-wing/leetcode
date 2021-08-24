package treesandgraphs

import (
	"container/list"
	"fmt"
)

func kthSmallest(root *TreeNode, k int) int {
	visitedNodes := make(map[int]bool)
	stack := list.New()
	stack.PushFront(root)
	index := 0
	var result int
	for nodeElement := stack.Front() ; stack.Len() > 0 ;{
		nodeElement = stack.Front()
		node := nodeElement.Value.(*TreeNode)
		if node.Left != nil && !visitedNodes[node.Left.Val]{
			stack.PushFront(node.Left)
			continue
		}
		visitedNodes[node.Val] = true
		index++
		if index == k {
			result = node.Val
			break
		}
		if node.Right != nil && !visitedNodes[node.Right.Val]{
			stack.PushFront(node.Right)
		}
		stack.Remove(nodeElement)
	}
	return result

}

func RunKthSmallest()  {
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
			},
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
			},
		},
		Right: &TreeNode{
			Val: 6,
		},
	}
	fmt.Printf("result:%v", kthSmallest(tree, 3))
}