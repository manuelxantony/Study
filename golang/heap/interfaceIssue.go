package main

import "container/heap"

type Node struct {
	Value int
}

type HeapTest []*Node

func testHeap() {
	nodeHeap := &HeapTest{}
	node := &Node{Value: 1}
	heap.Push(nodeHeap, node)
}

func (h HeapTest) Len() int {
	return len(h)
}

func (h HeapTest) Less(i, j int) bool {
	return h[i].Value < h[j].Value
}

func (h HeapTest) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *HeapTest) Push(val any) {
	*h = append(*h, val.(*Node))
}

func (h *HeapTest) Pop() any {
	temp := *h
	l := len(temp) - 1
	poped := temp[l]
	*h = temp[:l]
	return poped
}
