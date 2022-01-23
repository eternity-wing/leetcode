package problems

func twoSum(nums []int, target int) []int {
	var result []int
	numsMap := make(map[int]int, len(nums))

	for index, num := range nums {
		numsMap[num] = index
	}

	for index, num := range nums {
		lookup := target - num
		lookupIndex, ok := numsMap[lookup]
		if ok && lookupIndex != index {
			result = []int{lookupIndex, index}
			break
		}
	}

	return result
}
