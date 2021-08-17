package linkedlist

import (
	"math"
)

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	sizeOfA := 0
	sizeOfB := 0

	for aNode :=  headA ; aNode != nil && aNode.Next != nil; aNode = aNode.Next {
		sizeOfA++
	}

	for bNode :=  headB ; bNode != nil && bNode.Next != nil; bNode = bNode.Next {
		sizeOfB++
	}

	aNode := headA
	bNode := headB

	for i := 0 ; i < int(math.Abs(float64(sizeOfA-sizeOfB))) ; i++{
		if sizeOfA > sizeOfB{
			aNode = aNode.Next
		}else{
			bNode = bNode.Next
		}
	}

	for ; aNode != nil && bNode != nil ;{

		if aNode == bNode {
			return aNode
		}
		aNode = aNode.Next
		bNode = bNode.Next

	}

	return nil
}
