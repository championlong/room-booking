package main

import "sort"

func stoneGameVI(aliceValues []int, bobValues []int) int {
	n := len(aliceValues)
	total := make([]int, n)
	for i := 0; i < n; i++ {
		total[i] = aliceValues[i] + bobValues[i]
	}
	sort.Slice(total, func(i, j int) bool {
		return total[i] > total[j]
	})
	aTotal, bTotal := 0, 0
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			aTotal += total[i]
		}
		bTotal += bobValues[i]
	}
	diff := aTotal - bTotal
	if diff > 0 {
		return 1
	} else if diff == 0 {
		return 0
	} else {
		return -1
	}
}
