package problems

func firstUniqChar(s string) int {
	charsMap := make(map[int32]int, 26)
	for _, chr := range s {
		charsMap['a'-chr]++
	}

	for i, chr := range s {
		count, _ := charsMap['a'-chr]
		if count == 1 {
			return i
		}
	}
	return -1
}
