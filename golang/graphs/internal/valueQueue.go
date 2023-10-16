package internal

// one ndoe
type QueueNode struct {
	Value    string
	Distance int
}

type VDQueue []QueueNode

func (q *VDQueue) Enqueue(val ...QueueNode) {
	*q = append(*q, val...)
}

func (q *VDQueue) Dequeue() QueueNode {
	temp := *q
	dequeueVal := temp[0]
	*q = temp[1:]
	return dequeueVal
}

func (q VDQueue) IsEmpty() bool {
	return len(q) == 0
}
