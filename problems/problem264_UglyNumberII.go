package problems

import (
	"fmt"
	"leetcode/utils"
)

func nthUglyNumber(n int) int {
	nextTwoFactor := 0
	nextThreeFactor := 0
	nextFiveFactor := 0

	twoFactor := 2
	threeFactor := 3
	fiveFactor := 5

	memo := make([]int, n)
	memo[0] = 1
	for i := 1; i < n; i++ {
		memo[i] = utils.Min(utils.Min(twoFactor, threeFactor), fiveFactor)
		if memo[i] == twoFactor {
			nextTwoFactor += 1
			twoFactor = 2 * memo[nextTwoFactor]
		}
		if memo[i] == threeFactor {
			nextThreeFactor += 1
			threeFactor = 3 * memo[nextThreeFactor]
		}
		if memo[i] == fiveFactor {
			nextFiveFactor += 1
			fiveFactor = 5 * memo[nextFiveFactor]
		}
	}
	return memo[n-1]
}

func RunNthUglyNumber() {
	fmt.Printf("%v", nthUglyNumber(150))
}
