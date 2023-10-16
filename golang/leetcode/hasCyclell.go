package main


func hasCycle(head *ListNode) bool {
	mapNextRef := make(map [*ListNode]bool)

	for head != nil {
		mapNextRef[head] = true
		head = head.Next
		if _, ok := mapNextRef[head]; ok {
			return true
		}
	}

	return false
}