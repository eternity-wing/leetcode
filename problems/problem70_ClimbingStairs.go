package problems

import "fmt"

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	num1 := 1
	num2 := 2

	for num := 3; num < n; num++ {
		temSum := num1 + num2
		num1 = num2
		num2 = temSum
	}

	return num1 + num2
}

func RunClimbStairs() {
	fmt.Printf("%d", climbStairs(5))
}
