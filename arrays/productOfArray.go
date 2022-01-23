package arrays

import "fmt"

func productExceptSelf(nums []int) []int {
	product := 1
	numberOfZeros := 0
	for _, num := range nums {
		if num == 0 {
			numberOfZeros++
		} else {
			product *= num
		}
	}

	for index, num := range nums {
		if (num != 0 && numberOfZeros > 0) || (num == 0 && numberOfZeros > 1) {
			nums[index] = 0
		} else if num != 0 {
			nums[index] = product / num
		} else if numberOfZeros == 1 {
			nums[index] = product
		}

	}
	return nums
}

func RunProductExceptSelf() {
	fmt.Printf("%+v", productExceptSelf([]int{1,2,3,4}))
}
