package problems

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next

	for node := head; node != nil && node.Next != nil; {
		next := node.Next
		nextOfNext := next.Next

		next.Next = node
		node.Next = nextOfNext
		if nextOfNext != nil && nextOfNext.Next != nil {
			node.Next = nextOfNext.Next
		}

		node = nextOfNext
	}
	return newHead
}


func RunSwapPairs() {
	sw := swapPairs(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	})
	for node := sw; node != nil; node = node.Next {
		fmt.Printf("%d\t", node.Val)
	}

}
