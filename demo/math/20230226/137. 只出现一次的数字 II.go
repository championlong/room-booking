package main

func singleNumber(nums []int) int {
	countMap := make(map[int]int, 0)
	for _, x := range nums {
		countMap[x]++
	}
	for k, v := range countMap {
		if v == 1 {
			return k
		}
	}
	return -1
}
