package main

import (
	"fmt"
)

type Node struct {
	previous *Node
	data int // try any // generic data type
	next *Node
}

type List struct {
	head *Node
	tail *Node
}

// insert before
func (l *List) addBefore(newNode *Node, beforeNode *Node){
	prevNode := beforeNode.previous
	// insert before head
	if prevNode == nil{
		l.head = newNode
		l.head.next = beforeNode
		beforeNode.previous = l.head
		return
	}
	prevNode.next = newNode
	newNode.next = beforeNode
	beforeNode.previous = newNode
	newNode.previous = prevNode
}


// add front
func (l *List) addFront(value int) {
	newNode := &Node{data: value}

	// no nodes were present
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	l.addBefore(newNode, l.head)
}

// add after
func (l *List) addAfter(newNode *Node, afterNode *Node) {
	nextNode := afterNode.next
	// insert after tail
	if nextNode == nil {
		l.tail = newNode
		l.tail.previous = afterNode
		afterNode.next = l.tail
		return
	}

	afterNode.next = newNode
	newNode.next = nextNode
	nextNode.previous = newNode
	newNode.previous = afterNode
}

func (l *List) addEnd(value int) {
	newNode := &Node{data: value}
	
	if l.tail == nil {
		l.tail = newNode
		l.head = newNode
		return
	}

	l.addAfter(newNode, l.tail)
}

func (l *List) remove(value int) {
	if l.head == nil || l.tail == nil{
		return
	}

	// head node
	if l.head.data == value {
		nextNode := l.head.next
		l.head.next = nil
		nextNode.previous = nil
		l.head = nextNode
		return
	}

	// tail node
	if l.tail.data == value {
		previousNode := l.tail.previous
		l.tail.previous = nil
		previousNode.next = nil
		l.tail = previousNode
		return
	}

	// between node
	curr := l.head
	for curr.next != nil && curr.next.data != value {
		curr = curr.next
	}

	if curr.next != nil{
		curr.next = curr.next.next
		curr.next.previous = curr
	}

}

func (l *List) search(value int) (*Node, bool){
	curr := l.head
	for curr != nil && curr.data != value {
		curr = curr.next
	}
	if curr != nil {
		return curr, true
	}

	return nil, false
}

func (l *List) print() {
	// print head
	if v:= l.head; v != nil{
		 fmt.Println("head: " ,l.head.data)
	} else {
		fmt.Println("head nil")
	}
	// print tail
	 if v:= l.tail; v != nil{
	 	fmt.Println("tail: " ,l.tail.data)
	} else {
		fmt.Println("tail nil")
	}
	// print all node data
	curr := l.head
	for curr != nil {
		if curr.next != nil{
			fmt.Print(curr.data, "->")
		} else {
			fmt.Print(curr.data)
			fmt.Println()
		}
		curr = curr.next
	}
}

func main(){
	l := &List{}
	l.addEnd(1)
	l.addFront(2)
	l.addFront(3)
	l.addEnd(4)
	l.addEnd(5)
	l.addEnd(6)
	
	l.remove(6)
	l.print()

	if node, ok := l.search(6); ok {
		fmt.Println("value present: ", node)
	} else {
		fmt.Println("value not found")
	}
}