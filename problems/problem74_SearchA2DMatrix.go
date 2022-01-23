package problems

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	lowRow := 0
	highRow := m - 1

	for lowRow <= highRow {
		median := (lowRow + highRow) / 2

		if matrix[median][0] <= target && matrix[median][n-1] >= target {
			lowRow = median
			break
		}

		if matrix[median][0] < target {
			lowRow = median + 1
		} else {
			highRow = median - 1
		}
	}
	if lowRow >= m {
		return false
	}
	return binarySearch(target, matrix[lowRow])

}
func binarySearch(needle int, haystack []int) bool {

	low := 0
	high := len(haystack) - 1

	for low <= high {
		median := (low + high) / 2

		if haystack[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(haystack) || haystack[low] != needle {
		return false
	}

	return true
}

func RunSearchMatrix() {
	fmt.Printf("%v", searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 16))
}
