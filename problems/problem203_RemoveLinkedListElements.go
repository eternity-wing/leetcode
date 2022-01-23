package problems

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	node := head
	var previous *ListNode
	for node != nil {
		if node.Val == val {
			next := node.Next
			for ; next != nil && next.Val == val; next = next.Next {
			}

			if next != nil {
				node.Val = next.Val
				node.Next = next.Next
			} else if previous != nil {
				previous.Next = nil
			}else {
				return nil
			}
		}
		previous = node
		node = node.Next

	}
	return head
}

func RunRemoveElements() {
	tree := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
		},
	}

	node := removeElements(tree, 1)
	for node != nil {
		fmt.Printf("%d\t", node.Val)

		node = node.Next
	}

	//tree := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 1,
	//		Next: &ListNode{
	//			Val: 2,
	//			Next: &ListNode{
	//				Val: 3,
	//				Next: &ListNode{
	//					Val: 3,
	//				},
	//			},
	//		},
	//	},
	//}
}
