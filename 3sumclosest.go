package main

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	foundSum := nums[0] + nums[1] + nums[2]

	n := len(nums)

	for i := 0; i < n; i++ {
		j := i + 1
		k := len(nums) - 1

		for j < k {
			partialSum := nums[i] + nums[j] + nums[k]

			if abs(partialSum-target) < abs(foundSum-target) {
				foundSum = partialSum
			}

			if partialSum <= target {
				j++
			} else {
				k--
			}
		}

	}

	return foundSum
}

func abs(num int) int {
	if num < 0 {
		num = -num
	}

	return num
}
