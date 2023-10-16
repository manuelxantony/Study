package main

import (
	"fmt"
)

func (g *Graph) BFS(source string) {
	var queue Queue
	queue.Enqueue(source)

	for !queue.IsEmpty() {
		deq := queue.Dequeue()
		fmt.Println(deq)

		queue.Enqueue(g.Map[deq]...)
	}
}
