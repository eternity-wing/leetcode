package dynamicprogramming

import "fmt"

func uniquePaths(m int, n int) int {
	matrix := make([][]uint, m)
	for i := range matrix {
		matrix[i] = make([]uint, n)
	}

	for i := m - 1 ; i >= 0 ; i--{
		for j := n - 1 ; j >= 0 ; j--{
			pathThroughBottom := uint(0)
			if i+1 < m {
				pathThroughBottom = matrix[i+1][j]
			}

			pathThroughRight := uint(0)
			if j+1 < n {
				pathThroughRight = matrix[i][j+1]
			}
			numberOfPath := pathThroughBottom + pathThroughRight
			if numberOfPath == 0 {
				numberOfPath = 1
			}
			matrix[i][j] = numberOfPath
		}
	}
	return int(matrix[0][0])
}

func RunUniquePaths()  {
	fmt.Printf("%d", uniquePaths(2, 2))
}