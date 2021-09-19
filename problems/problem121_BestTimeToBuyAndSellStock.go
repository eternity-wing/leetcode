package problems

import "fmt"

func maxProfit(prices []int) int {
	dayWithStock := -1 * prices[0]
	dayWithOutStock := 0
	for i := 1; i < len(prices); i++ {
		price := prices[i]
		todayWithStock := maxInt(dayWithStock, dayWithOutStock-price)
		dayWithOutStock = maxInt(dayWithStock + price, dayWithOutStock)
		dayWithStock = todayWithStock
	}

	return dayWithOutStock
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func RunMaxProfit() {
	fmt.Printf("%d", maxProfit([]int{7,1,5,3,6,4}))
}
