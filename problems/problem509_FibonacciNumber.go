package problems

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	}
	previous := 0
	previousPrev := 1
	for i := 2; i <= n; i++ {
		temp := previous + previousPrev
		previousPrev = previous
		previous = temp
	}
	return previous + previousPrev
}

func RunFib()  {
	fmt.Printf("%d", fib(6))
}
