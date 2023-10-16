package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(num int) {
	q.items = append(q.items, num)
}

func (q *Queue) Dequeue() int {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) print(){
	for item := range q.items {
		fmt.Printf("%d ", item)
	}
}