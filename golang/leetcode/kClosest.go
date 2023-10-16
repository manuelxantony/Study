package main

import "container/heap"

type NodeG2 struct {
	Point    []int
	Distance int
}

type HeapNode2 []NodeG2

func kClosest2(points [][]int, k int) [][]int {

	nodeSlice := &HeapNode2{}

	for i := 0; i < len(points); i++ {
		point := points[i]
		distance := point[0]*point[0] + point[1]*point[1]
		node := NodeG2{
			Point:    point,
			Distance: distance,
		}
		heap.Push(nodeSlice, node)
	}

	out := [][]int{}
	for i := 0; i < k; i++ {
		heap.Init(nodeSlice)
		point := heap.Pop(nodeSlice)
		out = append(out, point.(NodeG2).Point)
	}

	return out
}

func (h HeapNode2) Len() int {
	return len(h)
}

// min pr max heap
func (h HeapNode2) Less(i, j int) bool {
	return h[i].Distance < h[j].Distance
}

func (h HeapNode2) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *HeapNode2) Push(val any) {
	*h = append(*h, val.(NodeG2))
}

func (h *HeapNode2) Pop() any {
	temp := *h
	l := len(temp) - 1
	poped := temp[l]
	*h = temp[:l]
	return poped
}
