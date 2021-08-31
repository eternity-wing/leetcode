package problems

import "fmt"

func maxSubArray(nums []int) int {
	max := -1000001
	sum := 0
	for _, num := range nums {
		if sum+num > num {
			sum += num
		} else {
			sum = num
		}

		if sum > max {
			max = sum
		}

	}
	return sum
}

func RunMaxSubArray() {
	fmt.Printf("%d", maxSubArray([]int{1}))
}
