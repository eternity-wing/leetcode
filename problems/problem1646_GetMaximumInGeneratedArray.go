package problems

import "fmt"

func getMaximumGenerated(n int) int {
	if n <= 1 {
		return n
	}
	results := make([]int, n+1)
	max := -1
	results[0] = 0
	results[1] = 1
	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			results[i] = results[i/2]
		} else {
			results[i] = results[i/2] + results[(i/2)+1]
		}
		if results[i] > max {
			max = results[i]
		}
	}
	return max
}

func RunGetMaximumGenerated() {
	fmt.Printf("%d", getMaximumGenerated(6))
}
