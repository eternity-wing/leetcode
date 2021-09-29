package problems

import "fmt"

func maxProfitII(prices []int) int {
	n := len(prices)

	if n <= 1 {
		return 0
	}

	twoDaysAgoWithOutStock := 0
	aDayAgoWithStock := -1 * prices[0]
	aDayAgoWithOutStock := 0


	for i := 1; i < len(prices); i++ {
		price := prices[i]


		todayWithStock := maxInt(aDayAgoWithStock, twoDaysAgoWithOutStock-price)
		todayWithOutStock := maxInt(aDayAgoWithOutStock, aDayAgoWithStock + price)

		twoDaysAgoWithOutStock = aDayAgoWithOutStock
		aDayAgoWithOutStock = todayWithOutStock
		aDayAgoWithStock = todayWithStock
	}
	return aDayAgoWithOutStock
}

func RunMaxProfitII()  {
	fmt.Printf("%d", maxProfitII([]int{1,2,3,0,2}))
}