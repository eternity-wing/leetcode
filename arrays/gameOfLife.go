package arrays


func gameOfLife(board [][]int) {

	rows := len(board)
	columns := len(board[0])
	copyBoard := make([][]int, rows)
	for index, _ := range copyBoard {
		copyBoard[index] = append([]int{}, board[index]...)
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < columns; col++ {
			neighborsState := getNeighborsState(copyBoard, rows, columns, row, col)
			isCellLive := copyBoard[row][col] == 1
			if isCellLive && (neighborsState.numberOfOnes < 2 || neighborsState.numberOfOnes > 3) {
				board[row][col] = 0
			}else if isCellLive && neighborsState.numberOfOnes <= 3{
				board[row][col] = 1
			}else if !isCellLive && neighborsState.numberOfOnes == 3{
				board[row][col] = 1
			}

		}
	}


}

type NeighborsState struct {
	numberOfZeros int
	numberOfOnes  int
}

func getNeighborsState(matrix [][]int, rows int, columns int, row int, col int) (result NeighborsState) {
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			neighborCol := col + j
			neighborRow := row + i
			if neighborRow == row && neighborCol == col {
				continue
			}
			if isValidCell(rows, columns, neighborRow, neighborCol) {
				result.numberOfOnes += matrix[neighborRow][neighborCol]
				result.numberOfZeros += matrix[neighborRow][neighborCol]
			}
		}
	}

	return result

}

func isValidCell(rows int, columns int, row int, col int) bool {
	if row < 0 || row >= rows {
		return false
	}
	if col < 0 || col >= columns {
		return false
	}
	return true
}

func RunGameOfLife() {
	gameOfLife([][]int{
		{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0},
	})
}
