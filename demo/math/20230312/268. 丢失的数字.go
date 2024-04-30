package main

func missingNumber(nums []int) int {
	numMap := make(map[int]bool, len(nums))
	for _, num := range nums {
		numMap[num] = true
	}
	for i := 0; i <= len(nums); i++ {
		if _, ok := numMap[i]; !ok {
			return i
		}
	}
	return len(nums)
}
