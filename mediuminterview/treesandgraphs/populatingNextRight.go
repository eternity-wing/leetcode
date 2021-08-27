package treesandgraphs

import (
	"container/list"
	"fmt"
)

func connect(root *Node) *Node {
	if root == nil{
		return root
	}
	currentLevelQueue := list.New()
	nextLevelQueue := list.New()

	currentLevelQueue.PushBack(root)
	for ; currentLevelQueue.Len() > 0 ; {
		topElement := currentLevelQueue.Front()
		topNode := topElement.Value.(*Node)
		if topNode.Left != nil{
			nextLevelQueue.PushBack(topNode.Left)
		}
		if topNode.Left != nil{
			nextLevelQueue.PushBack(topNode.Right)
		}
		if topElement.Next() != nil{
			topNode.Next = topElement.Next().Value.(*Node)
		}
		currentLevelQueue.Remove(topElement)
		if currentLevelQueue.Len() == 0 {
			currentLevelQueue = nextLevelQueue
			nextLevelQueue = list.New()
		}
	}

	return root
}


func RunConnect()  {
	root := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 4,
			},
			Right: &Node{
				Val: 5,
			},
		},
		Right: &Node{
			Val: 3,
			Left: &Node{
				Val: 6,
			},
			Right: &Node{
				Val: 7,
			},
		},
	}
	fmt.Printf("%+v", connect(root).Right.Left.Next)
}