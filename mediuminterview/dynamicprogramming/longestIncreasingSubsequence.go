package dynamicprogramming

import "fmt"


func lengthOfLIS(nums []int) int {
	max := 0
	n := len(nums)
	subArrayLength := make([]int, n)
	for i := 0; i < n ; i++{
		num := nums[i]
		subArrayLength[i] = 1
		maxI := 0
		for j := i-1; j >= 0 ; j--{
			if nums[j] < num && subArrayLength[j] > maxI{
				maxI = subArrayLength[j]
			}
		}
		subArrayLength[i] = maxI + 1
		max = maxInt(max, subArrayLength[i])
	}
	return max

}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}


func RunLengthOfLIS() {
	fmt.Printf("%d", lengthOfLIS([]int{10,9,2,5,3,7,101,18}))
}
