package problems

import (
	"fmt"
)

func nthSuperUglyNumber(n int, primes []int) int {
	primesIndexes := make([]int, len(primes))

	memo := make([]int, n)
	memo[0] = 1
	for i := 1; i < n; i++ {
		min := 9999999999
		var minIndexes []int
		for primeIndex, primeNumber := range primes {
			factor := primeNumber*memo[primesIndexes[primeIndex]]
			if min < factor {
				continue
			}
			if factor < min {
				minIndexes = []int{primeIndex}
			}else{
				minIndexes = append(minIndexes, primeIndex)
			}
			min = factor
		}
		memo[i] = min
		for _, minIndex := range minIndexes {
			primesIndexes[minIndex]++
		}
	}
	return memo[n-1]
}

func RunNthSuperUglyNumber() {
	fmt.Printf("%v", nthSuperUglyNumber(150, []int{2,3, 5}))
}
