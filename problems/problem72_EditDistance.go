package problems

import (
	"fmt"
	"leetcode/utils"
)

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	if m == 0 || n == 0 {
		return utils.Max(m, n)
	}

	resultsMap := make(map[int]map[int]int, m)
	for i := 0; i < m; i++ {
		resultsMap[i] = make(map[int]int, n)
	}

	var myFunc func(i int, j int) int

	myFunc = func(i int, j int) int {
		if i >= m {
			return n - j
		}

		if j >= n {
			return m - i
		}
		if _, ok := resultsMap[i][j]; ok {
			return resultsMap[i][j]
		}

		if word1[i] == word2[j] {
			resultsMap[i][j] = myFunc(i+1, j+1)
		} else {
			resultsMap[i][j] = 1 + utils.Min(utils.Min(myFunc(i+1, j+1), myFunc(i+1, j)), myFunc(i, j+1))
		}
		return resultsMap[i][j]
	}
	myFunc(0, 0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("[%d:%d]:%d\t", i, j, resultsMap[i][j])
		}
		fmt.Printf("\n")
	}
	return resultsMap[0][0]
}

func RunMinDistance() {
	fmt.Printf("%d", minDistance("bc", "ab"))
}
