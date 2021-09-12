package dynamicprogramming

import "fmt"

func lengthOfLIS(nums []int) int {
	max := 0
	n := len(nums)
	increasingSubsequencesMap := make(map[int]int)

	var myFunc func(value int, index int) int
	myFunc = func(value int, index int) int {
		for i := index ; i < n ; i++{
			if nums[i] > value{
				return increasingSubsequencesMap[i]
			}
		}
		return 0
	}

	for i := n - 1; i >= 0; i-- {
		increasingSubsequencesMap[i] = 1 + myFunc(nums[i], i+1)
		if increasingSubsequencesMap[i] > max {
			max = increasingSubsequencesMap[i]
		}
	}
	fmt.Printf("%+v\n", increasingSubsequencesMap)
	return max

}

func RunLengthOfLIS()  {
	fmt.Printf("%d", lengthOfLIS([]int{0,1,0,3,2,3}))
}
