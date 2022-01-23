package problems

import "fmt"

type DungeonPlayer struct {
	health int
	startHealth int
}

func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	n := len(dungeon[0])
	memo := make([]int, n)


	if dungeon[0][0] >= 0 {
		memo[0] = 1
	}else{
		memo[0] = dungeon[0][0] - 1
	}

	for j := 1; j < n; j++ {
		prev := memo[j-1]
		res := prev + dungeon[0][j]
		if prev >= 0 && res <= 0 {
			res = res - 1
		}
		memo[j] = res
	}
	fmt.Printf("%v", memo)
	return 1
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			previous := -2000
			if j-1 >= 0 {
				previous = memo[j-1]
			}
			res := dungeon[i][j] + maxInt(previous, memo[j])
			if res < 0{
				memo[j] = res
			}
		}
	}
	if memo[n-1] <= 0 {
		return (-1 * memo[n-1]) + 1
	}
	return 0
}

func RunCalculateMinimumHP() {
	fmt.Printf("%d", calculateMinimumHP([][]int{{-3, 5}}))
}


