package problems

import "fmt"

func matrixReshape(mat [][]int, r int, c int) [][]int {
	m := len(mat)
	n := len(mat[0])
	if r * c != m * n {
		return mat
	}

	result := make([][]int, r)
	for i := 0; i < r; i++ {
		result[i] = make([]int, c)
	}

	i1 := 0
	j1 := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j1 >= c {
				j1 = 0
				i1++
			}
			result[i1][j1] = mat[i][j]
			j1++
		}
	}
	return result
}

func RunMatrixReshape()  {
	fmt.Printf("%+v", matrixReshape([][]int{{1,2},{3,4}}, 2, 4))
}