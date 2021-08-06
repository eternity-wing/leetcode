package problems

import (
	"container/list"
	"fmt"
	"strconv"
)

type PuzzleNode struct {
	IOfZero int
	JOfZero int
	depth   int
	State   [][]int
}

func slidingPuzzle(board [][]int) int {
	i, j := getZeroPosition(board)
	puzzle := PuzzleNode{
		State:   board,
		depth:   0,
		JOfZero: j,
		IOfZero: i,
	}

	stateQueue := list.New()
	stateQueue.PushBack(puzzle)
	visitedStates := make(map[string]bool)

	goalStateStr := getStringPresentation([][]int{{1, 2, 3}, {4, 5, 0}})

	for frontElement := stateQueue.Front(); stateQueue.Len() != 0; frontElement = stateQueue.Front() {
		frontPuzzle := frontElement.Value.(PuzzleNode)
		stateQueue.Remove(frontElement)
		if goalStateStr == getStringPresentation(frontPuzzle.State) {
			return frontPuzzle.depth
		}
		prepareNextUnVisitedStates(frontPuzzle, stateQueue, visitedStates)
	}

	return -1
}

func getZeroPosition(board [][]int) (iOfZero int, jOfZero int) {

	for i := range board {
		for j := range board[i] {
			if board[i][j] == 0 {
				return i, j
			}
		}
	}

	return iOfZero, jOfZero
}

func getStringPresentation(board [][]int) string {
	presentation := ""
	for i := range board {
		for j := range board[i] {
			if presentation != "" {
				presentation += ","
			}
			presentation += strconv.Itoa(board[i][j])
		}
	}
	return presentation
}

func prepareNextUnVisitedStates(puzzle PuzzleNode, visitQueue *list.List, visitedStates map[string]bool) {
	for _, nextPuzzle := range puzzle.getNextPossibleStates() {
		stringPresentation := getStringPresentation(nextPuzzle.State)
		if !visitedStates[stringPresentation] {
			visitQueue.PushBack(nextPuzzle)
			visitedStates[stringPresentation] = true
		}
	}
}

func (p *PuzzleNode) getNextPossibleStates() (results []PuzzleNode) {

	tmp1 := make([][]int, 2)
	tmp1[0] = append(tmp1[0], p.State[0]...)
	tmp1[1] = append(tmp1[1], p.State[1]...)
	tmp1[0][p.JOfZero] = p.State[1][p.JOfZero]
	tmp1[1][p.JOfZero] = p.State[0][p.JOfZero]
	results = append(results, PuzzleNode{
		State:   tmp1,
		depth: p.depth + 1,
		JOfZero: p.JOfZero,
		IOfZero: 1 - p.IOfZero,
	})

	if p.JOfZero > 0 {
		tmp2 := make([][]int, 2)
		tmp2[0] = append(tmp2[0], p.State[0]...)
		tmp2[1] = append(tmp2[1], p.State[1]...)
		tmp2[p.IOfZero][p.JOfZero-1] = p.State[p.IOfZero][p.JOfZero]
		tmp2[p.IOfZero][p.JOfZero] = p.State[p.IOfZero][p.JOfZero-1]
		results = append(results, PuzzleNode{
			State:   tmp2,
			depth: p.depth + 1,
			JOfZero: p.JOfZero - 1,
			IOfZero: p.IOfZero,
		})
	}

	if p.JOfZero < 2 {
		tmp3 := make([][]int, 2)
		tmp3[0] = append(tmp3[0], p.State[0]...)
		tmp3[1] = append(tmp3[1], p.State[1]...)
		copy(tmp3[0], p.State[0])
		copy(tmp3[1], p.State[1])
		tmp3[p.IOfZero][p.JOfZero+1] = p.State[p.IOfZero][p.JOfZero]
		tmp3[p.IOfZero][p.JOfZero] = p.State[p.IOfZero][p.JOfZero+1]
		results = append(results, PuzzleNode{
			State:   tmp3,
			depth: p.depth + 1,
			JOfZero: p.JOfZero + 1,
			IOfZero: p.IOfZero,
		})
	}
	return results
}

func RunSlidingPuzzle() {
	fmt.Printf("\n%+v", slidingPuzzle([][]int{{3, 2, 4}, {1, 5, 0}}))
}
