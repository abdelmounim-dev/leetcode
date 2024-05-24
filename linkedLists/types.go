package linkedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func append(list *ListNode, node *ListNode) {
	for list.Next != nil {
		list = list.Next
	}
	list.Next = node
}
