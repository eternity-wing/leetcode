package linkedlist

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	node := head
	evenHead := head.Next
	i := 1
	oddTail := head
	for ; ; {
		if i%2 != 0 {
			oddTail = node
		}
		i++
		next := node.Next
		if next != nil{
			node.Next = next.Next
			node = next
		}else{
			break
		}
	}
	if oddTail != nil && evenHead != nil{
		oddTail.Next = evenHead
	}
	return head
}
