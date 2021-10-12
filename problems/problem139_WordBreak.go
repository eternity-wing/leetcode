package problems

import (
	"fmt"
	"strings"
)

func wordBreak(s string, wordDict []string) bool {
	strLen := len(s)
	memo := make(map[int]bool, strLen)

	var wordBreakInRange func (start int)bool
	wordBreakInRange = func(start int) bool {
		if start >= strLen {
			return true
		}
		if found, ok := memo[start] ; ok{
			return found
		}
		result := false
		tail := s[start:]
		for _, word := range wordDict {
			isStartWithWord := strings.HasPrefix(tail, word)
			if s == word || (isStartWithWord && wordBreakInRange(start + len(word))){
				result = true
				break
			}
		}
		memo[start] = result
		return memo[start]
	}
	wordBreakInRange(0)
	 return memo[0]
}

func RunWordBreak()  {
	fmt.Printf("%v", wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"}))
}
