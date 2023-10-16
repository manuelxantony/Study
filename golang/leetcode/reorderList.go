package main

import "fmt"

func reorderList(head *ListNode)  {
	if head != nil {
		nodeMap  := make(map[int]*ListNode)
 		i := 0
		curr := head

		for curr != nil {
			nodeMap[i] = curr
			curr = curr.Next
			i ++
		}
		i = i-1

		curr = head
		for j:=0; j < i/2; j++ {
			fmt.Println(nodeMap)
			next := curr.Next
			lastNode := nodeMap[i-j]
			curr.Next = lastNode
			lastNode.Next = next
			lastPrevNode := nodeMap[i-1-j]
			lastPrevNode.Next = nil
			// curr.Next.Next will never be nill as loop only goes half way
			curr = curr.Next.Next 
		}
	}
}