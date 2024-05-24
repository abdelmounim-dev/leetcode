package linkedlists

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	pointerMap := make(map[*ListNode]int)

	for node := head; node != nil; node = node.Next {
		if _, ok := pointerMap[node]; ok {
			return true
		}
		pointerMap[node] = node.Val
	}

	return false
}
