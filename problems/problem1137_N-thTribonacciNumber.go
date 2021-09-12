package problems

import "fmt"

func tribonacci(n int) int {
	if n <= 1 {
		return n
	} else if n == 2{
		return 1
	}
	first := 0
	second := 1
	third := 1

	for i := 3; i < n; i++ {
		temp := first + second + third
		first = second
		second = third
		third = temp
	}
	return first + second + third

}

func RunTribonacci()  {
	fmt.Printf("%d", tribonacci(2))
}
