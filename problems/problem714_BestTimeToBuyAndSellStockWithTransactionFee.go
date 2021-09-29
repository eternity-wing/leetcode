package problems

import "fmt"

func maxProfitTransactional(prices []int, fee int) int {
	dayWithStock := -1 * prices[0]
	dayWithOutStock := 0
	for i := 1; i < len(prices); i++ {
		price := prices[i]
		todayWithStock := maxInt(dayWithStock, dayWithOutStock-price)
		dayWithOutStock = maxInt(dayWithOutStock, dayWithStock + price - fee)
		dayWithStock = todayWithStock
	}

	return dayWithOutStock
}


func RunMaxProfitTransactional() {
	fmt.Printf("%d", maxProfitTransactional([]int{1,3,7,5,10,3}, 3))
}
