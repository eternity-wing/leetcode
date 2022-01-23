package problems

import (
	"container/list"
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil{
		return result
	}
	myQueue := list.List{}
	childrenQueue := list.List{}
	myQueue.PushFront(root)
	frontElement := myQueue.Front()
	var levelItems []int
	for frontElement != nil {

		frontNode := frontElement.Value.(*TreeNode)
		levelItems = append(levelItems, frontNode.Val)
		if frontNode.Left != nil {
			childrenQueue.PushBack(frontNode.Left)
		}
		if frontNode.Right != nil {
			childrenQueue.PushBack(frontNode.Right)
		}

		frontElement = frontElement.Next()
		if frontElement == nil && childrenQueue.Len() > 0 {
			result = append(result, levelItems)
			myQueue = childrenQueue
			childrenQueue = list.List{}
			levelItems = []int{}
			frontElement = myQueue.Front()
		}

	}
	if len(levelItems) > 0{
		result = append(result, levelItems)
	}
	return result
}

func RunLevelOrder() {
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
	fmt.Printf("%v", levelOrder(root))
}
