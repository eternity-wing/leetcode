package problems

import "fmt"

func trap(height []int) int {
	startPointer := 0
	endPointer := 1
	n := len(height)
	totalRain := 0

	filledInRange := 0
	for endPointer < n {
		startHeight := height[startPointer]
		endHeight := height[endPointer]
		if endHeight >= startHeight || endPointer == n - 1{
			totalRangeCapacity := (1 + endPointer - startPointer) * startHeight
			startPointer = endPointer
			endPointer = endPointer + 1
			totalRain += totalRangeCapacity - filledInRange
			filledInRange = 0
			continue
		}else{
			filledInRange += height[endPointer]
			endPointer++
		}


	}

	return totalRain
}

func RunTrap()  {
	fmt.Printf("%d", trap([]int{0,1,0,2,1,0,1,3,2}))
}