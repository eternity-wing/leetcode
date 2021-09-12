package problems

import (
	"fmt"
	"leetcode/utils"
)

func minCostClimbingStairs(cost []int) int {
	numberOfStairs := len(cost)

	for i := numberOfStairs - 1; i >= 0; i-- {
		nextCost := 0
		nextOfNextCost := 0

		if i+1 < numberOfStairs {
			nextCost = cost[i+1]
		}
		if i+2 < numberOfStairs {
			nextOfNextCost = cost[i+2]
		}

		cost[i] = cost[i] + utils.Min(nextCost, nextOfNextCost)
	}

	return utils.Min(cost[0], cost[1])

}

func RunMinCostClimbingStairs()  {
	fmt.Printf("%d", minCostClimbingStairs([]int{10,15,20}))
}
