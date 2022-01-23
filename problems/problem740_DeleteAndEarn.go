package problems

import "fmt"

func deleteAndEarn(nums []int) int {
	//incidents := make([]int, 10001)
	//
	//uniqueNumbersCount := 0
	//for _, num := range nums{
	//	if incidents[num] == 0{
	//		uniqueNumbersCount++
	//	}
	//	incidents[num] += num
	//}
	//
	//max := -10000
	//
	//results := make([][]int, 2)
	//results[0] = make([]int, uniqueNumbersCount)
	//results[1] = make([]int, uniqueNumbersCount)
	//
	//i := 0
	//for num, sum := range incidents{
	//	if sum == 0 {
	//		continue
	//	}
	//}
	//
	//
	//for num, _ := range incidents{
	//	prevCount := 0
	//	if _, ok := incidents[num-2] ; ok{
	//		prevCount = incidents[num-2]
	//	}
	//	incidents[num] += prevCount
	//	max = maxInt(max, incidents[num])
	//}
	//fmt.Printf("%v\n", incidents)

	return 10

}

func RunDeleteAndEarn()  {
	fmt.Printf("%d", deleteAndEarn([]int{1,1,1,2,4,5,5,5,6}))
}