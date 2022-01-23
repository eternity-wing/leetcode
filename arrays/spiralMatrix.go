package arrays

import "fmt"

const (
	SpiralRight = iota
	SpiralDown
	SpiralLeft
	SpiralUp
)



func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	result := make([]int, m*n)
	i := 0
	j := 0
	x := 0
	const VisitedCellValue = 200
	direction := SpiralRight
	for ; x < (m * n); {
		result[x] = matrix[i][j]
		x++
		matrix[i][j] = VisitedCellValue
		switch direction {
		case SpiralRight:
			if j+1 < n && matrix[i][j+1] != VisitedCellValue {
				j++
			}else{
				direction = SpiralDown
				i++
			}
			break
		case SpiralDown:
			if i+1 < m && matrix[i+1][j] != VisitedCellValue {
				i++
			}else{
				direction = SpiralLeft
				j--
			}
			break
		case SpiralLeft:
			if j-1 >= 0 && matrix[i][j-1] != VisitedCellValue {
				j--
			}else{
				direction = SpiralUp
				i--
			}
			break
		case SpiralUp:
			if i-1 >= 0 && matrix[i-1][j] != VisitedCellValue {
				i--
			}else{
				direction = SpiralRight
				j++
			}
			break
		}
	}

	return result
}

func RunSpiralOrder() {
	//expected[1,2,3,6,9,8,7,4,5]
	//fmt.Printf("%v", spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))

	//[1,2,3,4,8,12,16,15,14,13,9,5,6,7,11,10]
	fmt.Printf("%v", spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}))
}
