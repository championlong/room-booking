package main

func main() {
	swapPairs(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	})
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1->2->3->4
// 2->1->4->3
func swapPairs(head *ListNode) *ListNode {
	result := &ListNode{}
	result.Next = head // 0->1->2->3->4

	prev := result // 0->1->2->3->4
	for head != nil && head.Next != nil {
		//0->2->3->4
		prev.Next = head.Next
		// 3->4
		next := head.Next.Next
		// 1->2->1->2->...
		head.Next.Next = head
		// 1->3->4
		head.Next = next
		// 1->3->4
		prev = head
		// 3->4
		head = next
	}
	return result.Next
}
