package problems

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	previous := head

	node := head
	for node != nil {
		if node.Val == previous.Val && node.Next == nil {
			previous.Next = nil
		} else if node.Val != previous.Val {
			previous.Next = node
			previous = node
		}

		node = node.Next
	}
	return head
}

func RunDeleteDuplicates() {
	tree := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},
	}
	node := deleteDuplicates(tree)
	for node != nil {
		fmt.Printf("%d\t", node.Val)

		node = node.Next
	}
}
