package problems

import (
	"fmt"
	"leetcode/utils"
)

func intersect(nums1 []int, nums2 []int) []int {
	numsMap := make(map[int][]int)
	var result []int
	for _, num := range nums1{
		if _, ok := numsMap[num] ; !ok{
			numsMap[num] = []int{0, 0}
		}
		numsMap[num][0]++
	}
	for _, num := range nums2{
		if _, ok := numsMap[num] ; !ok{
			numsMap[num] = []int{0, 0}
		}
		numsMap[num][1]++
	}
	
	for num, memo := range numsMap{
		for i := 0 ; i < utils.Min(memo[0], memo[1]) ; i++{
			result = append(result, num)
		}
	}
	return result
}

func RunIntersect()  {
	fmt.Printf("%v", intersect([]int{9,4,9,8,4}, []int{4, 9, 5}))
}