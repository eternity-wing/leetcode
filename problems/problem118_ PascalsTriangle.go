package problems

import "fmt"

func PascalsTriangle(numRows int) [][]int {
	results := make([][]int, numRows)
	results[0] = []int{1}
	for i := 1; i < numRows; i++ {
		results[i] = generateNextPascalTriangleRow(results[i-1], i)
	}
	return results
}
func generateNextPascalTriangleRow(row []int, rowNumber int) []int {
	nextRowNumber := rowNumber + 1
	nextRow := make([]int, nextRowNumber)
	for i := 0; i < nextRowNumber; i++ {
		upperLeftNumber := 0
		if i - 1 >= 0 {
			upperLeftNumber = row[i-1]
		}
		upperNumber := 0
		if i  < rowNumber {
			upperNumber = row[i]
		}

		nextRow[i] = upperLeftNumber + upperNumber
	}
	return nextRow
}

func RunGeneratePascalsTriangle()  {
	fmt.Printf("%+v", PascalsTriangle(3))
}
