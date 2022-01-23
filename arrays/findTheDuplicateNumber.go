package arrays

import "fmt"

func findDuplicate(nums []int) int {
	sum := 0
	n := len(nums) - 1

	for _, num := range nums {
		sum += num
	}
	return sum - ((n * (n + 1)) / 2)
}

func RunFindDuplicate()  {
	fmt.Printf("%d", findDuplicate([]int{3,1,3,4,2}))
}