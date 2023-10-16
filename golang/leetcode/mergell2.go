package main

import (
	"sort"
)

func mergeKLists2(lists []*ListNode) *ListNode {
	nodeValues := []int{}
	for _, list := range lists {
		curr := list
		for curr != nil {
			nodeValues = append(nodeValues, curr.Val)
			curr = curr.Next
		}
	}
	if len(nodeValues) == 0 {
		return nil
	}
	sort.Ints(nodeValues)

	lNode := &ListNode{Val: nodeValues[0]}
	head := lNode
	for _, nodeVal := range nodeValues[1:] {
		lNode.Next = &ListNode{Val: nodeVal}
		lNode = lNode.Next
	}
	return head
}
