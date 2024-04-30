package main

func isPalindrome(head *ListNode) bool {
	nums := []int{}
	for cur := head; cur != nil; cur = cur.Next {
		nums = append(nums, cur.Val)
	}
	for left, right := 0, len(nums)-1; left <= right; {
		if nums[left] != nums[right] {
			return false
		}
		left++
		right--
	}
	return true
}
