package problems

import (
	"fmt"
	"leetcode/utils"
)

func trap(height []int) int {
	n := len(height)
	right := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		right[i] = maxInt(right[i+1], height[i])
	}

	left := make([]int, n)
	left[0] = height[0]
	for i := 1; i < n; i++ {
		left[i] = maxInt(left[i-1], height[i])
	}


	totalRain := 0


	for i := 1; i < n-1; i++ {
		totalRain += utils.Min(left[i], right[i]) - height[i]
	}

	return totalRain
}

func RunTrap() {
	//fmt.Printf("%d", trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2}))
	fmt.Printf("%d", trap([]int{4,2,0,3,2,5}))
}
