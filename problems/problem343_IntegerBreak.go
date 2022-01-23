package problems

import "fmt"

func integerBreak(n int) int {
	products := make(map[int]int, n)
	products[1] = 1
	products[2] = 1

	var myFunc func(num int) int

	myFunc = func(num int) int {
		if num <= 0 {
			return 0
		}
		if _, ok := products[num]; ok {
			return products[num]
		}

		max := num - 1
		for i := 2; i <= (n/2) + 1; i++ {
			max = maxInt(i * myFunc(num-i) , max)
			max = maxInt(i * (num-i) , max)
		}
		products[num] = max
		return products[num]
	}
	return myFunc(n)
}

func RunIntegerBreak() {
	fmt.Printf("%d", integerBreak(5))
}
