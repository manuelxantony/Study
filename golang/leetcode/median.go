package main

import "container/heap"

type Heap struct {
	arr  []int
	less func(i, j int) bool
}

type MedianFinder struct {
	minHeap *Heap
	maxHeap *Heap
}

func (h Heap) Len() int { return len(h.arr) }

func (h *Heap) Push(val any) { h.arr = append(h.arr, val.(int)) }

func (h *Heap) Pop() any {
	l := len(h.arr) - 1
	poped := h.arr[l]
	h.arr = h.arr[:l]
	return poped
}

func (h Heap) Less(i, j int) bool { return h.less(h.arr[i], h.arr[j]) }

func (h *Heap) Swap(i, j int) { h.arr[i], h.arr[j] = h.arr[j], h.arr[i] }

// to see the first value
func (h *Heap) Peek() int {
	peekedVal := h.arr[0]
	return peekedVal
}

func Constructor1() MedianFinder {
	return MedianFinder{
		maxHeap: &Heap{
			// arr: []int{},
			less: func(i, j int) bool {
				return i > j
			},
		},
		minHeap: &Heap{
			// arr: []int{},
			less: func(i, j int) bool {
				return i < j
			},
		},
	}
}

// idea is to have a maxheap and minheap
// maxheap hold values lesser than minHeap
// thus we are creating sorted values for popping in both heaps
// and those values will be median values

// example:
// arr := 4,51,20, 7, 1, 30
// maxheap = 7, 1, 4
// 			7
//		   / \
//        1   4
// minheap = 20, 51, 30
// 			20
//		   / \
//        51   30
// here 7 and 20 will be in sorted order of the array

//  1. len of maxheap and minheap should only differ by 1(approx equal)
//  2. every value in maxheap should be lesser than minheap
//  3. to ensure that peek from maxHeap and minHeap and
//  4. if the peek value from maxHeap is larger than peek  value in  maxheap
//  5. pop from maxheap and push the value to minheap thus criteria 2 will be satisfy
//  6. for finding median ther are 2 cases
//     case 1. if len(maxheap) > len(minheap)
//     then pop from maxheap, that is median
//     case 2. if len(maxheap) < len(minheap)
//     then pop from minheap, that is median
//     case 3. if len(maxheap) == len(minheap)
//     then (pop from maxheap + pop from minheap) / 2 , that is median

func (this *MedianFinder) AddNum(num int) {
	if this.maxHeap.Len() == 0 {
		heap.Push(this.maxHeap, num)
	} else {
		if num < this.maxHeap.Peek() {
			heap.Push(this.maxHeap, num)
		} else {
			heap.Push(this.minHeap, num)
		}
	}
	if (this.maxHeap.Len() - this.minHeap.Len()) > 1 {
		poped := heap.Pop(this.maxHeap)
		heap.Push(this.minHeap, poped)
	} else if (this.minHeap.Len() - this.maxHeap.Len()) > 1 {
		poped := heap.Pop(this.minHeap)
		heap.Push(this.maxHeap, poped)
	}

}

func (this *MedianFinder) FindMedian() float64 {
	var median float64
	if this.maxHeap.Len() == this.minHeap.Len() {
		median = (float64(this.minHeap.Peek()) + float64(this.maxHeap.Peek())) / 2.0
	} else if this.maxHeap.Len() > this.minHeap.Len() {
		median = float64(this.maxHeap.Peek())
	} else {
		median = float64(this.minHeap.Peek())
	}
	return median
}
