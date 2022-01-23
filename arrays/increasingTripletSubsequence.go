package arrays

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	small := math.Inf(0)
	big := math.Inf(0)
	for _, num := range nums {
		floatNum := float64(num)
		if floatNum <= small {
			small = floatNum
		} else if floatNum <= big {
			big = floatNum
		} else {
			return true
		}
	}
	return false

}



//func increasingTriplet(nums []int) bool {
//	n := len(nums)
//	maxes := make([]int, n)
//	mins := make([]int, n)
//	maxes[n-1] = nums[n-1]
//	mins[0] = nums[0]
//
//	for i := n - 2; i >= 0; i-- {
//		if nums[i] >= maxes[i+1] {
//			maxes[i] = nums[i]
//		} else {
//			maxes[i] = maxes[i+1]
//		}
//	}
//
//	for i := 1; i < n; i++ {
//		if nums[i] < mins[i-1] {
//			mins[i] = nums[i]
//		} else {
//			mins[i] = mins[i-1]
//		}
//	}
//
//	for index, num := range nums {
//		if num > mins[index] && num < maxes[index]{
//			return true
//		}
//	}
//	return false
//
//}

func RunIncreasingTriplet() {
	//fmt.Printf("%v", increasingTriplet([]int{20, 100, 10, 12, 5, 13}))
	//fmt.Printf("%v", increasingTriplet([]int{1,2,1,3}))
	//fmt.Printf("%v", increasingTriplet([]int{1, 2, 3, 4, 5}))
	fmt.Printf("%v", increasingTriplet([]int{5, 4, 3, 2, 1}))
}
