package internal

type Queue []string

func (q *Queue) Enqueue(val ...string) {
	*q = append(*q, val...)
}

func (q *Queue) Dequeue() string {
	temp := *q
	dequeueVal := temp[0]
	*q = temp[1:]
	return dequeueVal
}

func (q Queue) IsEmpty() bool {
	return len(q) == 0
}
