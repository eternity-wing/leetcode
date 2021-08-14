package linkedlist

import "fmt"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sum *ListNode
	sumTail := sum
	carry := 0

	firstNumber := l1
	secondNumber := l2
	
	for ;firstNumber != nil || secondNumber != nil;{
		firstDigit := 0
		secondDigit := 0
		if firstNumber != nil{
			firstDigit = firstNumber.Val
		}
		if secondNumber != nil{
			secondDigit = secondNumber.Val
		}
		
		digitsSum := firstDigit + secondDigit + carry
		if sumTail != nil{
			sumTail.Next = &ListNode{Val: digitsSum % 10}
			sumTail = sumTail.Next
		}else{
			sum = &ListNode{Val: digitsSum % 10}
			sumTail = sum
		}

		carry = digitsSum / 10
		if firstNumber!= nil{
			firstNumber = firstNumber.Next
		}

		if secondNumber!= nil{
			secondNumber = secondNumber.Next
		}
	}

	if carry > 0 && sumTail != nil {
		sumTail.Next = &ListNode{Val: carry}
	}
	
	return sum
	
}

func RunAddTwoNumbers()  {
	first := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	second := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	fmt.Printf("\n%+v", addTwoNumbers(first, second))

}