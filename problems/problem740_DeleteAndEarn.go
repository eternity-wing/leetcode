package problems

import "fmt"

func deleteAndEarn(nums []int) int {
	incidents := make(map[int]int)
	totalAmount := 0

	for _, num := range nums{
		incidents[num] += num
		totalAmount += num
	}

	max := -10000

	for num, _ := range incidents{
		max = maxInt(max, totalAmount - (incidents[num-1] + incidents[num+1]))
	}

	return max

}

func RunDeleteAndEarn()  {
	fmt.Printf("%d", deleteAndEarn([]int{}))
}