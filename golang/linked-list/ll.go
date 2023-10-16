package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node
	length int

}

func (l *List) add(value int) {
	newNode := &Node{data: value}
	l.length ++

	if l.head == nil {
		l.head = newNode
		
		return
	}

	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode
}


func (l *List) addSorted(value int) {
	newNode := &Node{data: value}
	l.length ++

	if l.head == nil {
		l.head = newNode
		return
	}

	if l.head.data > value {
		curr := l.head
		l.head = newNode
		l.head.next = curr
		return
	}

	curr := l.head
	for curr.next != nil && curr.next.data < value {
		curr = curr.next
	}

	biggerNode := curr.next
	curr.next = newNode
	newNode.next = biggerNode
}


func (l *List) remove(value int) {
	if l.head == nil {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length --
		return
	}

	curr := l.head 

	for curr.next != nil && curr.next.data != value {
		curr = curr.next
	}
	
	if curr.next != nil {
		curr.next = curr.next.next
		l.length --
	}
}

func (l *List) print() {
	curr := l.head
	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.next
	}
}