package problems

import "fmt"

func numTrees(n int) int {
	memo := make([][]int, n)
	for row := 0; row < n; row++ {
		memo[row] = make([]int, n)
	}

	var numTreesInRange func(start int, end int) int
	numTreesInRange = func(start int, end int) int {
		if start >= end {
			return 1
		}

		if memo[start-1][end-1] > 0 {
			return memo[start-1][end-1]
		}

		result := 0

		for i := start; i <= end; i++ {
			leftChoices := numTreesInRange(start, i-1)
			rightChoices := numTreesInRange(i+1, end)
			result += leftChoices * rightChoices
		}
		memo[start-1][end-1] = result
		return result
	}
	s := numTreesInRange(1, n)

	return s
}



func RunNumTrees() {
	fmt.Printf("%d", numTrees(4))
}
