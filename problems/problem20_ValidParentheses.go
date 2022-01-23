package problems

import (
	"container/list"
)

func isValid(s string) bool {
	myStack1 := list.List{}
	myStack2 := list.List{}
	myStack3 := list.List{}

	for _, chr := range s {
		if chr == '}' || chr == '{' {
			if chr == '}' {
				if !checkStack(myStack1, '{') {
					return false
				} else {
					myStack1.Remove(myStack1.Front())
				}
			} else {
				myStack1.PushFront(chr)
			}
		} else if chr == ')' || chr == '(' {
			if chr == ')' {
				if !checkStack(myStack2, '(') {
					return false
				} else {
					myStack2.Remove(myStack2.Front())
				}
			} else {
				myStack2.PushFront(chr)
			}
		} else {
			if chr == ']' {
				if !checkStack(myStack3, '[') {
					return false
				} else {
					myStack3.Remove(myStack3.Front())
				}
			} else {
				myStack3.PushFront(chr)
			}
		}
	}
	return myStack1.Front() == nil && myStack2.Front() == nil && myStack3.Front() == nil
}

func checkStack(stack list.List, char int32) bool {
	front := stack.Front()
	if front == nil || front.Value != char {
		return false
	}
	return true
}
