package arrays


func RemoveDuplicates(nums []int) int {
	uniqueValuePointer := 0
	previousVal := -101

	n := len(nums)

	for i := 0 ; i < n ; i++{
		num := nums[i]
		if previousVal == num {
			continue
		}
		previousVal= num
		nums[uniqueValuePointer] = num
		uniqueValuePointer++
	}


	return uniqueValuePointer
}