package _0221219

import "math"

func increasingTriplet(nums []int) bool {
	n := len(nums)
	first := nums[0]
	second := math.MaxInt32
	for i := 1; i < n; i++ {
		if nums[i] > second {
			return true
		}
		if nums[i] < second && nums[i] > first {
			second = nums[i]
		}
		if nums[i] < first {
			first = nums[i]
		}
	}
	return false
}
