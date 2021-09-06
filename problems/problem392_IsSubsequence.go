package problems

import "fmt"

func isSubsequence(s string, t string) bool {
	lengthOfS := len(s)
	i := 0

	if s == "" {
		return true
	}

	for j, _ := range t {
		if t[j] == s[i] {
			i++
		}

		if lengthOfS == i {
			return true
		}
	}
	return false
}

func RunIsSubsequence() {
	fmt.Printf("%v", isSubsequence("abc", "ahbgdc"))
}
