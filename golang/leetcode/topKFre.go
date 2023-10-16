package main

import (
	"container/heap"
	"fmt"
)

type CountNode struct {
	count int
	value int
}

type Count []*CountNode

func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]*CountNode)
	countOut := &Count{}
	for i := 0; i < len(nums); i++ {
		if node, ok := countMap[nums[i]]; ok {
			node.count++
		} else {
			node := &CountNode{
				count: 1,
				value: nums[i],
			}
			countMap[nums[i]] = node
		}
	}

	for _, valueNode := range countMap {
		fmt.Println(valueNode)
		heap.Push(countOut, valueNode)
	}

	out := make([]int, k)
	for i := 0; i < k; i++ {
		out[i] = heap.Pop(countOut).(CountNode).value
	}

	return out
}

func (n Count) Len() int { return len(n) }

// max heap
func (n Count) Less(i, j int) bool { return n[i].count > n[j].count }

func (n Count) Swap(i, j int) { n[i], n[j] = n[j], n[i] }

func (n *Count) Push(x any) {
	*n = append(*n, x.(*CountNode))
}

func (n *Count) Pop() any {
	old := *n
	l := len(old) - 1
	poped := old[l]
	*n = old[:l]
	return poped
}
