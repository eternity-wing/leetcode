package problems

import (
	"fmt"
	"leetcode/utils"
)

func minimumTotal(triangle [][]int) int {
	min := triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		min = 1000000
		for j, num := range triangle[i] {
			topRightIndex := j
			if topRightIndex >= len(triangle[i-1]) {
				topRightIndex = j - 1
			}

			topLeftIndex := j - 1
			if topLeftIndex < 0 {
				topLeftIndex = topRightIndex
			}

			triangle[i][j] = num + utils.Min(triangle[i-1][topRightIndex], triangle[i-1][topLeftIndex])
			min = utils.Min(min, triangle[i][j])
		}
	}
	return min
}

func RunMinimumTotal() {
	//fmt.Printf("%d", minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
	fmt.Printf("%d", minimumTotal([][]int{{-11}}))
}
