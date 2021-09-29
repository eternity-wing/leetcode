package problems

import "fmt"

func numberOfArithmeticSlices(nums []int) int {
	const outOfRangeNumber = -2001
	startSequenceIndex := 0
	sequenceDiff := outOfRangeNumber
	numberOfSlices := 0
	n := len(nums)
	cnt := 0
	for i := 0; i < n; i++ {
		nextItem := outOfRangeNumber
		if i+1 < n{
			nextItem = nums[i+1]
		}

		sliceLength := i + 1 - startSequenceIndex
		if sliceLength >= 3 {
			cnt += sliceLength - 2
		}else{
			cnt = 0
		}
		diff := nextItem - nums[i]
		if diff == sequenceDiff {
			continue
		}

		sequenceDiff = diff
		startSequenceIndex = i
		numberOfSlices += cnt
	}

	return numberOfSlices
}

func RunNumberOfArithmeticSlices() {
	fmt.Printf("%d", numberOfArithmeticSlices([]int{1, 2, 4, 5, 6}))
}
