package treesandgraphs

import (
	"fmt"
)

func numIslands(grid [][]byte) int {
	numberOfRows := len(grid)
	numberOfColumns := len(grid[0])
	numberOfIslands := 0


	var doDfs = func(row int, column int) {}
	doDfs = func(row int, column int) {
		grid[row][column] = '0'
		neighbors := getNeighbors(row, column, numberOfRows, numberOfColumns)
		for _, neighbor := range neighbors {
			neighborRow := neighbor[0]
			neighborColumn := neighbor[1]
			if grid[neighborRow][neighborColumn] == '1' {
				doDfs(neighborRow, neighborColumn)
			}
		}
	}

	for row := 0; row < numberOfRows; row++ {
		for column := 0; column < numberOfColumns; column++ {
			cellVal := grid[row][column]
			if cellVal == '0' {
				continue
			}
			numberOfIslands++
			doDfs(row, column)
		}
	}
	return numberOfIslands
}

func getNeighbors(row int, column int, numberOfRows int, numberOfColumns int) (neighbors [][]int) {
	if row-1 >= 0 {
		neighbors = append(neighbors, []int{row - 1, column})
	}
	if row+1 < numberOfRows {
		neighbors = append(neighbors, []int{row + 1, column})
	}
	if column-1 >= 0 {
		neighbors = append(neighbors, []int{row, column - 1})
	}
	if column+1 < numberOfColumns {
		neighbors = append(neighbors, []int{row, column + 1})
	}
	return neighbors
}

func getCellKey(row int, column int) string {
	return fmt.Sprintf("%d-%d", row, column)
}


func RunNumIslands() {
	//[["1","1","0","0","0"],["1","1","0","0","0"],["0","0","1","0","0"],["0","0","0","1","1"]]
	//fmt.Printf("%d", numIslands([][]byte{
	//	{'1', '1', '0', '0', '0'},
	//	{'1', '1', '0', '0', '0'},
	//	{'0', '0', '1', '0', '0'},
	//	{'0', '0', '0', '1', '1'}}))
	fmt.Printf("%d", numIslands([][]byte{
		{'1', '1', '1', '1', '1'}}))
}
