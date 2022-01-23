package problems

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	node := head
	var previous *ListNode

	for node != nil {
		next := node.Next
		node.Next = previous
		previous = node
		node = next
	}

	return previous
}

func RunReverseList()  {
	tree := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	node := reverseList(tree)
	for node != nil {
		fmt.Printf("%d\t", node.Val)

		node = node.Next
	}
}