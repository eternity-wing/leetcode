package problems

import (
	"fmt"
	"leetcode/utils"
)

func jump(nums []int) int {
	aBigNumber := 10000000
	n := len(nums)
	lastIndex := n-1
	nums[lastIndex] = 0
	for i := lastIndex - 1 ; i >= 0 ; i--{
		val := nums[i]
		longestAccessibleIndex := i + val
		if longestAccessibleIndex >= lastIndex {
			nums[i] = 1
			continue
		}

		min := aBigNumber
		for j := i+1 ; j <= longestAccessibleIndex ; j++{
			min = utils.Min(min, nums[j])
		}

		nums[i] = 1 + min

	}

	return nums[0]

}

func RunCanJumpII()  {
	fmt.Printf("%v\n", jump([]int{1}))
}