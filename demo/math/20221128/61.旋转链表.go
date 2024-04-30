package main

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	l := 1
	cur := head
	for cur.Next != nil {
		cur = cur.Next
		l++
	}

	k %= l
	cur.Next = head
	for i := l - k; i > 0; i-- {
		cur = cur.Next
	}

	head = cur.Next
	cur.Next = nil
	return head
}
