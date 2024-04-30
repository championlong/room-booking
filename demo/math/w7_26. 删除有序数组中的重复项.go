package main

import "fmt"

func main() {
	removeDuplicates([]int{0, 1, 1, 1, 1, 2, 2, 3, 3, 4})
}

func removeDuplicates(nums []int) int {
	j := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[j] = nums[i]
			j++
		}
	}
	fmt.Println(nums[:j])
	return j
}

//
//func removeDuplicates(nums []int) int {
//	tmpMap := make(map[int]int, len(nums))
//	for i := 0; i < len(nums); i++ {
//		if _, ok := tmpMap[nums[i]]; ok {
//			nums = append(nums[:i], nums[i+1:]...)
//			i--
//		} else {
//			tmpMap[nums[i]] = nums[i]
//		}
//	}
//	return len(nums)
//}
