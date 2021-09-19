package problems

import "fmt"

func rob(nums []int) int {
	max := 0
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		noneAdjacentMax := 0
		if i+2 < n {
			noneAdjacentMax = nums[i+2]
		}
		res := nums[i] + noneAdjacentMax
		if res > max {
			max = res
		}
		nums[i] = max
	}
	return max
}
func RunRob() {
	fmt.Printf("%d", rob([]int{1, 3, 1}))
}
