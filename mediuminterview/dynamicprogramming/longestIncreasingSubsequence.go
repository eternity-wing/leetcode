package dynamicprogramming

import "fmt"

type increasingSequence struct {
	head int
	size int
}

func lengthOfLIS(nums []int) int {
	max := 0
	n := len(nums)
	var increasingSubsequences []increasingSequence

	for i := n - 1; i >= 0; i-- {
		num := nums[i]
		newSequences := []increasingSequence{{head: num, size: 1}}
		for _, sequence := range increasingSubsequences {
			if sequence.head > num {
				newSequence := increasingSequence{head: num, size: sequence.size + 1}
				newSequences = append(newSequences, newSequence)
				if newSequence.size > max {
					max = newSequence.size
				}
			}
		}
		increasingSubsequences = append(increasingSubsequences, newSequences...)
		if max == 0 {
			max = 1
		}

	}
	return max

}

func RunLengthOfLIS() {
	fmt.Printf("%d", lengthOfLIS([]int{7,7,7,7,7,7,7}))
}
