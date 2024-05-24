package linkedlists

func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	reversedList := reverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return reversedList
}
