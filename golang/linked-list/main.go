package main

import "fmt"

// import "fmt"

// // type Node struct {
// // 	data int
// // 	next *Node
// // }

// // type List struct {
// // 	head *Node
// // }

// // func (list *List) add(value int){
// // 	newNode := &Node{data: value}

// // 	if list.head == nil {
// // 		list.head = newNode
// // 		return
// // 	}

// // 	curr := list.head
// // 	for curr.next != nil {
// // 		curr = curr.next
// // 	}

// // 	curr.next = newNode

// // }

// // sorted add
// func (list *List) addSorted(value int) {
// 	newNode := &Node{data: value}

// 	if list.head == nil {
// 		list.head = newNode
// 		return
// 	}

// 	curr := list.head
// 	if curr.data > value {
// 		list.head = newNode
// 		newNode.next = curr
// 		return
// 	}

// 	for curr.next != nil && curr.next.data < value {
// 		curr = curr.next
// 	}

// 	node := curr.next
// 	curr.next = newNode
// 	newNode.next = node
// }

// func (list *List) remove(value int){
// 	if list.head == nil {
// 		return
// 	}

// 	if list.head.data == value {
// 		list.head = list.head.next
// 		return
// 	}

// 	curr := list.head
// 	for curr.next != nil && curr.next.data != value{
// 		curr = curr.next
// 	}

// 	if curr.next != nil {
// 		curr.next = curr.next.next
// 	}

// }

// func (list *List) print() {
// 	curr := list.head
// 	for curr != nil {
// 		fmt.Println(curr.data)
// 		curr = curr.next
// 	}
// }

func main(){
	// list := &List{}
	// list.addSorted(10)
	// list.addSorted(2)
	// list.addSorted(33)
	// list.addSorted(4)

	// list.remove(15)
	// list.remove(2)

	// list.addSorted(1)

	// list.print()

	list := &List{}
	
	list.addSorted(4)
	list.addSorted(3)
	list.addSorted(5)
	list.print()
	fmt.Println("length: " ,list.length)

	list.remove(5)
	list.print()
	fmt.Println("length: ",list.length)

	list.remove(3)
	list.remove(4)
	list.print()
	fmt.Println("length: ",list.length)

	list.remove(10)
	list.print()
	fmt.Println("length: ",list.length)


	
}

