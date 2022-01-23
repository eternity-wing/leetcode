package problems

import "fmt"

func isValidSudoku(board [][]byte) bool {
	checkMap := make(map[string]bool)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			val := int(board[i][j] - '0')

			squareNumber := ((i / 3) * 3) + (j / 3) + 1

			bigSquareRow := fmt.Sprintf("0-0-%d-%d", i, val)
			bigSquareColumn := fmt.Sprintf("0-1-%d-%d", j, val)

			if _, ok := checkMap[bigSquareRow]; ok {
				return false
			}
			if _, ok := checkMap[bigSquareColumn]; ok {
				return false
			}
			smallSquareRow := fmt.Sprintf("%d-0-%d-%d", squareNumber, i%3, val)
			smallSquareColumn := fmt.Sprintf("%d-1-%d-%d", squareNumber, j%3, val)
			if _, ok := checkMap[smallSquareRow]; ok {
				return false
			}
			if _, ok := checkMap[smallSquareColumn]; ok {
				return false
			}

			checkMap[bigSquareRow] = true
			checkMap[bigSquareColumn] = true
			checkMap[smallSquareRow] = true
			checkMap[smallSquareColumn] = true
		}
	}

	return true

}

func RunIsValidSudoku() {

	fmt.Printf("%v\n", isValidSudoku([][]byte{
		{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
		{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
		{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
		{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
		{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
		{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
		{'.', '.', '4', '.', '.', '.', '.', '.', '.'},
	}))

	//fmt.Printf("%v\n", isValidSudoku([][]byte{
	//{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
	//{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	//{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	//{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	//{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	//{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	//{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	//{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	//{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	//}))

	//fmt.Printf("%v\n", isValidSudoku([][]byte{
	//	{'8','3','.','.','7','.','.','.','.'},
	//	{'6','.','.','1','9','5','.','.','.'},
	//	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	//	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	//	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	//	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	//	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	//	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	//	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	//}))

}
