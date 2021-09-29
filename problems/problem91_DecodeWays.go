package problems

import "fmt"

const NumberOfAlphabet = 26

//func numDecodings(s string) int {
//	n := len(s)
//	memo := make(map[int]int, n)
//
//	var decodeInRange func(index int) int
//	decodeInRange = func(index int) int {
//		if index < 0 || index >= n {
//			return 1
//		}
//		if found, ok := memo[index]; ok {
//			return found
//		}
//		number := convertRuneNumberToInt(s[index])
//		currentCases := 0
//		nextNumber := NumberOfAlphabet + 1
//		if index+1 < n {
//			nextNumber = convertRuneNumberToInt(s[index+1])
//		}
//		isInRange := (number*10)+nextNumber <= NumberOfAlphabet
//		if number == 0 {
//			currentCases = 0
//		} else if isInRange {
//			currentCases = decodeInRange(index+1) + decodeInRange(index+2)
//		} else {
//			currentCases = decodeInRange(index + 1)
//		}
//		memo[index] = currentCases
//		return currentCases
//
//	}
//
//	return decodeInRange(0)
//
//}

func numDecodings(s string) int {
	n := len(s)
	nextCases := 1
	nextOfNextCases := 1
	for i := n - 1; i >= 0; i-- {
		number := convertRuneNumberToInt(s[i])
		previousNumber := 3
		if i-1 >= 0 {
			previousNumber = convertRuneNumberToInt(s[i-1])
		}
		if number == 0 && (previousNumber >= 3 || previousNumber == 0) {
			return 0
		}

		nextNumber := NumberOfAlphabet + 1

		currentCases := 1
		if i+1 < n {
			nextNumber = convertRuneNumberToInt(s[i+1])
		}


		isInRange := (number*10)+nextNumber <= NumberOfAlphabet

		if number == 0 {
			nextOfNextCases = nextCases
			nextCases = 0
			continue
		} else if isInRange {
			currentCases = nextCases + nextOfNextCases
		} else{
			currentCases = nextCases
		}

		nextOfNextCases = nextCases
		nextCases = currentCases
	}

	return nextCases

}

func convertRuneNumberToInt(numb uint8) int {
	return int(numb - '0')
}

func RunNumDecodings() {
	fmt.Printf("%d", numDecodings("226"))
}
