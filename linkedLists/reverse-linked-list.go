package linkedlists

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var p, n *ListNode
	for c := head; c != nil; c = n {
		n = c.Next
		c.Next = p
		p = c
	}

	return p
}
