package problems

import "fmt"

func trap(height []int) int {
	filledSpaces := 0
	maxHeightIndex := 0
	totalRain := 0

	for index, h := range height{
		if height[maxHeightIndex] <= h{

			length := index - maxHeightIndex
			if length > 1{
				totalRain += (length - 1) * height[maxHeightIndex]
			}
			filledSpaces = 0
			maxHeightIndex = index
		}else {
			filledSpaces += h
		}
	}

	return totalRain
}

func RunTrap()  {
	fmt.Printf("%d", trap([]int{4,2,0,3,2,5}))
}