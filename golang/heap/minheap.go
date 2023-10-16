package main

type Heap []int


func CreateMinHeapFromArray(arr []int) *Heap {
	h := make(Heap, 0)
	for _, ele := range arr {
		h.push(ele)
	}
	return &h
}

func (h *Heap) push(val int) {
	if len(*h) == 0{
		*h = append(*h, val)
		return	
	}
	// append to the end
	*h = append(*h, val)
	h.heapyifyUp()	
}

func (h *Heap) pop() {
	temp := *h
	if len(temp) == 0{return}
	*h = temp[:len(temp)-1]
	h.heapyifyDown()
}

func (h Heap) heapyifyUp() {
	currentIndex := len(h)-1
	for h.hasParent(currentIndex) && currentIndex > 0{
		parentIndex := h.getParentIndex(currentIndex)
		if h[currentIndex] < h[parentIndex]{
			h[currentIndex], h[parentIndex] = h[parentIndex], h[currentIndex]
			currentIndex, parentIndex = parentIndex, currentIndex
		} else{
			break
		}
	}
}

func (h Heap) heapyifyDown(){
	currentIndex := 0
	for (h.hasLeftChild(currentIndex)){
		// if both left and right are present
		smallerIndex := h.getLeftChildIndex(currentIndex)
		if h.hasRightChild(currentIndex) && (h[smallerIndex] > h[h.getRightChildIndex(currentIndex)]) {
				smallerIndex = h.getRightChildIndex(currentIndex)
		}
		if h[currentIndex] < h[smallerIndex]{
			break
		} else{
			//swap
			h[currentIndex], h[smallerIndex] = h[smallerIndex], h[currentIndex]
			currentIndex = smallerIndex
		}
	}
}

func (h Heap) getParentIndex(childIndex int) int{ return (childIndex - 1) / 2 }

func (h Heap) getLeftChildIndex(parentIndex int) int{ return (parentIndex *2 + 1)}

func (h Heap) getRightChildIndex(parentIndex int) int{ return (parentIndex * 2 + 2)}

func (h Heap) hasParent(childIndex int) bool { return h.getParentIndex(childIndex) >= 0}

func (h Heap) hasLeftChild(parentIndex int) bool{ return h.getLeftChildIndex(parentIndex) < len(h)}

func (h Heap) hasRightChild(parentIndex int) bool{ return h.getRightChildIndex(parentIndex) < len(h)}