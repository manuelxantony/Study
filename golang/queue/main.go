package main

import "fmt"

//type New func(int)

type Queue struct {
	queue 	[]int	
}

func New() *Queue{
	return &Queue{}
}

func (q *Queue) enqueue(value int) {
	q.queue = append(q.queue, value)
}

func (q *Queue) dequeue() int {
	if len(q.queue) <= 0 {
		return -1
	}
	value := q.queue[0]
	fmt.Println("Deqeued: ", value)
	q.queue = q.queue[1:]
	return value
}

func (q *Queue) print() {
	for _, value := range q.queue{
		fmt.Print(value, " ")
	}
}
 
func main() {
	// q := New()
	// q.enqueue(1)
	// q.enqueue(2)
	// q.enqueue(3)
	// q.enqueue(4)
	// q.dequeue() // 1
	// q.dequeue() // 2
	// q.dequeue() // 3
	// q.print()
	// q.dequeue() // 4
	// q.print()
	// q.dequeue()
	// q.dequeue()
	// q.print()
	// aMain()
	aMainStack()
}