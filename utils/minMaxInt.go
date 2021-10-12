package utils

func Min(num1 int, num2 int) int {
	if num1 <= num2{
		return num1
	}
	return num2
}
func Max(num1 int, num2 int) int {
	if num1 >= num2{
		return num1
	}
	return num2
}

func MinOfArray(nums []int) int {
	min := nums[0]
	for _, num := range nums{
		min = Min(num, min)
	}
	return min
}