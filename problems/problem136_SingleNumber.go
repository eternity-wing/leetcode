package problems

func singleNumber(nums []int) int {
	var result int
	n := len(nums)
	numbersMap := make(map[int]int, (n/2) + 1)
	for _, num := range nums {
		numbersMap[num]++
	}

	for num, repeatCount := range numbersMap {
		if repeatCount == 1 {
			result = num
			break
		}
	}
	return result
}

