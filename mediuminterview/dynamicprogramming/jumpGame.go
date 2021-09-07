package dynamicprogramming

import "fmt"


func canJump(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return true
	}

	lastIndex := n - 1
	memory := make(map[int]bool, n)
	memory[lastIndex] = true

	for i := lastIndex - 1; i >= 0; i-- {
		memory[i] = false
		if i+nums[i] >= lastIndex {
			memory[i] = true
			continue
		}
		for j := i + 1; j <= lastIndex && j <= i+nums[i]; j++ {
			if memory[j] {
				memory[i] = true
				break
			}
		}

	}

	return memory[0]

}

func RunCanJump() {
	fmt.Printf("%+v", canJump([]int{}))
}
