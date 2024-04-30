package main

func findLucky(arr []int) int {
	tempMap := map[int]int{}
	for _, v := range arr {
		tempMap[v]++
	}

	max := -1
	for k, v := range tempMap {
		if k == v && k > max {
			max = k
		}
	}
	return max
}
