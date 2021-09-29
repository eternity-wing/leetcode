package problems

import "fmt"

func robII(nums []int) int {
	max := 0
	n := len(nums)
	maxes := make([]int, n)
	maxes[n-1] = n-1
	for i := n - 1; i >= 0; i-- {
		noneAdjacentMax := 0
		if i+2 < n && nums[i+2] > nums[maxes[i+2]]{
			noneAdjacentMax = nums[i+2]
		}else if i+2 < n{
			noneAdjacentMax = nums[maxes[i+2]]
		}
		res := nums[i] + noneAdjacentMax
		if res > max {
			max = res
		}
		nums[i] = max
	}
	return max
}
func RunRobII() {
	fmt.Printf("%d", rob([]int{1, 2, 3}))
}
