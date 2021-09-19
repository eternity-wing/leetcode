package dynamicprogramming

import "fmt"

func longestPalindrome(s string) string {
	maxLength := 1
	strLen := len(s)
	startIndex := 0

	resultMap := make([][]bool, strLen)
	for index, _ := range resultMap {
		resultMap[index] = make([]bool, strLen)
	}

	for i := 0; i < strLen; i++ {
		resultMap[i][i] = true
	}

	for i := 0; i < strLen-1; i++ {
		if s[i] == s[i+1] {
			resultMap[i][i+1] = true
			maxLength = 2
			startIndex = i
		}
	}

	for k := 3; k <= strLen; k++ {
		for i:=0 ; i < strLen - k + 1; i++{
			j := i + k - 1

			if resultMap[i+1][j-1] && s[i] == s[j]{
				resultMap[i][j] = true
				if k > maxLength{
					maxLength = k
					startIndex = i
				}
			}

		}
	}

	return s[startIndex:startIndex+maxLength]
}

func RunLongestPalindrome()  {
	fmt.Printf("%s", longestPalindrome("babad"))
}
