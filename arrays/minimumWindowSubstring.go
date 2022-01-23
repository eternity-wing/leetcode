package arrays

import "fmt"


func minWindow(s string, t string) string {
	alphabetSize := 'z' - 'A'
	sLength := len(s)
	tLength := len(t)
	tChars := make([]int, alphabetSize)
	windowChars := make([]int, alphabetSize)

	for _, chr := range t {
		tChars[getCharOffset(uint8(chr))]++
	}

	firstMatchedIndex := 0
	for ; firstMatchedIndex < sLength && tChars[getCharOffset(s[firstMatchedIndex])] == 0; firstMatchedIndex++ {
	}

	substrStart := -1
	substrLength := 1000000
	frequencyMatches := 0
	index := firstMatchedIndex
	windowStart := firstMatchedIndex
	windowEnd := firstMatchedIndex

	for ;; {

		charOffset := getCharOffset(s[index])
		windowChars[charOffset]++

		if windowChars[charOffset] == tChars[charOffset] {
			frequencyMatches += tChars[charOffset]
		}
		if frequencyMatches == tLength {
			matchedSubstrLength := index - firstMatchedIndex
			if matchedSubstrLength < substrLength {
				substrLength = matchedSubstrLength
				substrStart = firstMatchedIndex
			}
			frequencyMatches -= tChars[getCharOffset(s[firstMatchedIndex])]
			windowChars[getCharOffset(s[firstMatchedIndex])]--
			i := firstMatchedIndex + 1
			for ; tChars[getCharOffset(s[i])] == 0 && i < index; i++ {
				windowChars[getCharOffset(s[i])]--
			}
			firstMatchedIndex = i
			continue
		}

		index++
	}

	if substrStart == -1 {
		return ""
	}
	return s[substrStart : substrStart+substrLength+1]
}

func getCharOffset(chr uint8) int {
	return int(chr - 'A')
}

func RunMinWindow() {
	fmt.Printf("%s", minWindow("a", "b"))
}
