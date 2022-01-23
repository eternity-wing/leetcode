package arrays

import (
	"container/list"
	"fmt"
)

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)

	if n < k {
		return []int{}
	}

	result := make([]int, n + 1 -k)
	deq := list.List{}

	var prepareQue func(i int)

	prepareQue = func(i int) {
		if deq.Len() != 0 && (i - k) == deq.Front().Value.(int) {
			deq.Remove(deq.Front())
		}

		for deq.Back() != nil && nums[i] > nums[deq.Back().Value.(int)] {
			deq.Remove(deq.Back())
		}

	}

	for i := 0; i < k; i++ {
		prepareQue(i)
		deq.PushBack(i)
	}
	result[0] = nums[deq.Front().Value.(int)]
	j := 1

	for i := k; i < len(nums); i++ {
		prepareQue(i)
		deq.PushBack(i)
		result[j] = nums[deq.Front().Value.(int)]
		j++
	}

	return result
}

func RunMaxSlidingWindow() {
	fmt.Printf("%v", maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
