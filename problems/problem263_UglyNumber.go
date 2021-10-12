package problems

import "fmt"

func isUgly(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	primes := []int{2, 3, 5}

	for _, prime := range primes {
		if n%prime == 0 && isUgly(n/prime) {
			return true
		}
	}

	return false
}

func RunIsUgly()  {
	fmt.Printf("%v", isUgly(5))
}