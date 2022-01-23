package problems

func containsDuplicate(nums []int) bool {
	numsMap := make(map[int]bool , len(nums))
	for _, num :=range nums{
		if _, ok := numsMap[num] ; ok {
			return true
		}
		numsMap[num] = true
	}
	return false
}