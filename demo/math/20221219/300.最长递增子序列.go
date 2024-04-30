package _0221219

import "sort"

func lengthOfLIS(nums []int) int {
	result := make([]int, 0)
	for i := range nums {
		index := sort.Search(len(result), func(j int) bool {
			return result[j] >= nums[i]
		})
		if index == len(result) {
			result = append(result, nums[i])
		} else {
			result[index] = nums[i]
		}
	}
	return len(result)
}
