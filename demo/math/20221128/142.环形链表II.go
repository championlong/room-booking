package main

func detectCycle(head *ListNode) *ListNode {
	var null *ListNode
	if head == nil || head.Next == nil {
		return null
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if slow != fast {
		return null
	}
	fast = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return fast
}
