package problems

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	j := n - 1
	i := m - 1
	for index := m + n - 1; index >= m; index-- {
		max := 0
		if j >= 0 && nums2[j] >= nums1[i] {
			max = nums2[j]
			j--
		} else {
			max = nums1[i]
			i--
		}
		nums1[index] = max
	}

	fmt.Printf("%v\n%v", nums1, nums2)

}

func RunMerge() {
	merge([]int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3)
}
