package problems

import "fmt"

func canJump(nums []int) bool {
	n := len(nums)
	memo := make(map[int]bool, n)
	var recursiveCanJump func(index int) bool
	recursiveCanJump = func(index int) bool {
		if index >= n-1{
			return true
		}
		if found, ok := memo[index] ; ok{
			return found
		}
		memo[index] = false
		for i:=index ; i <= index + nums[index] && i < n ; i++{
			if recursiveCanJump(i){
				memo[index] = true
				break
			}
		}
		return memo[index]
	}

	return recursiveCanJump(0)
}

func RunCanJump()  {
	fmt.Printf("%v\n", canJump([]int{0}))
}