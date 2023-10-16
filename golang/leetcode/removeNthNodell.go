package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    nodeMap := make(map[int]*ListNode)
	curr := head
	i := 0
	for curr != nil {
		nodeMap[i] = curr
		curr = curr.Next
		i++
	}

	if n == i {
		head = head.Next
		return head
	}

	valuToRemove := i - n
	nodeToRemove := nodeMap[valuToRemove]

	nextNode := nodeToRemove.Next
	preNode := nodeMap[valuToRemove - 1]
	preNode.Next = nextNode
	return head
}