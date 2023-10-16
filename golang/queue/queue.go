package main

import (
	"container/list"
	"fmt"
)

func aMain(){
	   // new linked list
    queue := list.New()

    // Simply append to enqueue.
    queue.PushBack(10)
    queue.PushBack(20)
    queue.PushBack(30)

    // Dequeue
    front:=queue.Front()
    fmt.Println(front.Value)
    // This will remove the allocated memory and avoid memory leaks
    queue.Remove(front)
}