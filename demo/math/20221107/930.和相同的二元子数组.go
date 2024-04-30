package main

import (
	"fmt"
)

func main() {
	//nums := []int{1, 0, 1, 0, 1}
	nums := []int{0,0,0,0,0}
	result := numSubarraysWithSum2(nums, 0)
	fmt.Println(result)
}

/*
给你一个二元数组 nums ，和一个整数 goal ，请你统计并返回有多少个和为 goal 的 非空 子数组。

子数组 是数组的一段连续部分。
 */

// 时间超时
func numSubarraysWithSum1(nums []int, goal int) int {
	var result int
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == goal {
				result++
			}
		}
	}
	return result
}

func numSubarraysWithSum2(nums []int, goal int) int {
	return slidingWindow(nums, goal) - slidingWindow(nums, goal-1)
}

//【1, 0, 1, 0, 1】
// <= 2 - <= 1
//[1]
//[0] [1,0]
//[1,0,1] [0,1] [1]
//[1,0,1,0] [0,1,0] [1,0] [0]
//[0,1,0,1] [1,0,1] [0,1] [1]
func slidingWindow(nums []int, goal int) int {
	if goal < 0 {
		return 0
	}
	var result, sum, left, right int
	for right < len(nums) {
		sum += nums[right]
		for sum > goal {
			sum -= nums[left]
			left++
		}
		result += right - left + 1
		right++
	}
	return result
}

