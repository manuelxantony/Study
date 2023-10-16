package main

type goHeap []int

func (h goHeap) Len() int {
	return len(h)
}

func (h goHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h goHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *goHeap) Push(val any) {
	*h = append(*h, val.(int))
}

func (h *goHeap) Pop() any {
	temp := *h
	l := len(temp) - 1
	poped := temp[l]
	*h = temp[:l]
	return poped
}
