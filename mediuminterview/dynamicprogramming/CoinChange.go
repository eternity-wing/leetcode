package dynamicprogramming

import (
	"fmt"
	"math"
)

const COIN_CHANGE_MAX_AMOUNT = 10000

func coinChange(coins []int, amount int) int {
	calculatedAmounts := make(map[int]int)

	var coinChange1 func(amnt int) int
	coinChange1 = func(amnt int) int {
		if result, ok := calculatedAmounts[amnt] ; ok {
			return result
		}
		if amnt == 0 {
			return 0
		}
		min := COIN_CHANGE_MAX_AMOUNT + 1
		for _, coin := range coins {
			if coin > amnt {
				continue
			}
			res, ok := calculatedAmounts[amnt-coin]
			if !ok {
				res = coinChange1(amnt-coin)
			}
			if res < 0 {
				continue
			}
			min = int(math.Min(float64(1+res), float64(min)))
		}

		calculatedAmounts[amnt] = -1
		if min != COIN_CHANGE_MAX_AMOUNT+1 {
			calculatedAmounts[amnt] = min
		}

		return calculatedAmounts[amnt]

	}

	return coinChange1(amount)

}

func RunCoinChange() {
	fmt.Printf("%d", coinChange([]int{1}, 2))
}
