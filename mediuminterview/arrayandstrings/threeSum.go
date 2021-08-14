package arrayandstrings

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var result [][]int

	numsLength := len(nums)

	if numsLength < 3 {
		return result
	}
	sort.Ints(nums)

	for i := 0; i < numsLength; i++ {
		firstNum := nums[i]
		if i != 0 && firstNum == nums[i-1] {
			continue
		}

		for j, k := i+1, numsLength-1; j < k; {
			sum := firstNum + nums[j] + nums[k]
			if sum == 0 {
				result = append(result, []int{firstNum, nums[j], nums[k]})
				x := j
				for ; x < numsLength && nums[j] == nums[x]; x++ {
				}
				j = x
			} else if sum > 0 {
				k--
			} else {
				j++
			}
		}
	}
	return result
}

func RunThreeSum() {
	fmt.Printf("result:%v", threeSum([]int{0, 0, 0}))
}
