package main

import "sort"

func findContentChildren(g []int, s []int) int {

	var result int

	sort.Ints(g)
	sort.Ints(s)

	i, j := 0, 0

	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			result++
			i++
		}

		j++
	}

	return result

}
