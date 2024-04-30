package main

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	if prev != nil {
		prev.Next = nil
	}
	left := sortList(head)
	right := sortList(slow)
	return merge(left, right)
}

func merge(left, right *ListNode) *ListNode {
	temp := &ListNode{}
	cur := temp
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
		cur = cur.Next
	}
	if left != nil {
		cur.Next = left
	}
	if right != nil {
		cur.Next = right
	}
	return temp.Next
}
