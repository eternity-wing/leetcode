package problems

import "fmt"

func maxProfit(prices []int) int {
	MIN := 10001
	maxProfit := 0
	for _, num := range prices {
		if num-MIN > maxProfit {
			maxProfit = num - MIN
		}

		if num < MIN {
			MIN = num
		}
	}
	return maxProfit
}

func RunMaxProfit()  {
	fmt.Printf("%d", maxProfit([]int{7,6,4,3,1}))
}
