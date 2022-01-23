package problems

func canConstruct(ransomNote string, magazine string) bool {
	charsMap := make(map[rune]int)
	for _, chr := range magazine {
		if _, ok := charsMap[chr] ; !ok {
			charsMap[chr] = 0
		}
		charsMap[chr]++
	}

	for _, chr := range ransomNote {
		if count, ok := charsMap[chr] ; !ok || count <=0 {
			return false
		}
		charsMap[chr]--
	}

	return true
}
