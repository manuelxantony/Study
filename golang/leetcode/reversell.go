package main

type ListNode struct{
	Val int
	Next *ListNode
}

func reverseList(N *ListNode) *ListNode {
	head := N
	if head == nil {
		return head
	}

	pre := head
	next := head.Next
	pre.Next = nil

	for next != nil {
		head = next
		next = head.Next
		head.Next = pre
		pre = head
	}

	return head
}