package problems

import "fmt"

func PascalsTriangleII(rowIndex int) []int {

	result := make([]int, rowIndex + 1)
	result[0] = 1
	for i := 0; i < rowIndex + 1; i++ {
		for j := i; j  > 0; j-- {
			upperLeftNumber := 0
			if j - 1 >= 0 {
				upperLeftNumber = result[j - 1]
			}
			upperNumber := result[j]

			result[j] = upperLeftNumber + upperNumber
		}
	}
	return result
}
func RunGeneratePascalsTriangleII()  {
	fmt.Printf("%+v", PascalsTriangleII(4))
}