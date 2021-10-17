package problems

import "fmt"

func robII(nums []int) int {
	n := len(nums)
	if n == 1{
		return nums[0]
	}
	skipFirst := rob(append([]int{}, nums[1:]...))
	skipLast := rob(append([]int{}, nums[:n-1]...))

	return maxInt(skipFirst,skipLast)
}


func RunRobII() {
	fmt.Printf("%d", robII([]int{200, 3, 140, 20, 10}))
}
