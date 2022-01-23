package problems

import "fmt"

func isAnagram(s string, t string) bool {
	charsMap := make([]int, 26)
	for _, chr := range s {
		charsMap[chr - 'a']++
	}
	for _, chr := range t {
		charsMap[chr - 'a']++
	}

	for _, chr := range t {
		charsMap[chr - 'a']--
	}
	for _, chr := range s {
		charsMap[chr - 'a']--
	}

	for _, val := range charsMap{
		if val != 0 {
			return false
		}
	}

	return true
}

func RunIsAnagram()  {
	fmt.Printf("%v", isAnagram("anagram", "nagaram"))
}