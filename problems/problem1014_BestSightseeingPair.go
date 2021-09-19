package problems

import "fmt"

func maxScoreSightseeingPair(values []int) int {
	startIndex := 0
	max := 0
	n := len(values)
	for i := 1; i < n; i++ {
		currentValue := values[i]
		tmp := values[startIndex] + startIndex - i
		score := currentValue + tmp
		if score > max {
			max = score
		}
		
		if tmp < currentValue{
			startIndex = i
		}
	}
	return max

}

func RunMaxScoreSightseeingPair()  {
	fmt.Printf("%d", maxScoreSightseeingPair([]int{1, 2}))
}