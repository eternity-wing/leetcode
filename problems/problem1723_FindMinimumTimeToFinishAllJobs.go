package problems

import (
	"fmt"
	"sort"
)

func minimumTimeRequired(jobs []int, k int) int {
	times := make([]int, k)
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i] < jobs[j]
	})


	for _, job := range jobs{
		minTimeIndex := 0
		for i, time := range times {
			if time  < times[minTimeIndex]{
				minTimeIndex = i
			}
		}
		times[minTimeIndex] += job
	}

	max := -1
	for _, time := range times {
		if time  > max{
			max = time
		}
	}
	fmt.Printf("%v\n", times)

	return max
}

func RunMinimumTimeRequired()  {
	fmt.Printf("%d", minimumTimeRequired([]int{1,2,4,7,8}, 2))
}