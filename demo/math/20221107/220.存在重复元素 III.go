package main

import (
	"fmt"
	"math"
)

func main() {
	//nums := []int{1, 2, 3, 1}
	nums := []int{4,1,-1,6,5}
	result := containsNearbyAlmostDuplicate(nums, 3, 1)
	fmt.Println(result)
}
/*
给你一个整数数组 nums 和两个整数 k 和 t 。
请你判断是否存在 两个不同下标 i 和 j，使得 abs(nums[i] - nums[j]) <= t ，同时又满足 abs(i - j) <= k 。
如果存在则返回 true，不存在返回 false。
 */
func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	for i := 0; i < len(nums); i++ {
		//两个不同下标 i 和 j
		for j := i + 1; j - i <= indexDiff; j++ {
			if j == len(nums) {
				break
			}
			if math.Abs(float64(nums[i]-nums[j])) <= float64(valueDiff) {
				return true
			}
		}
	}
	return false
}
