package main

import "fmt"

type Node struct {
	Value 		int
	Next 		*Node
	MinmumValue int
}

type MinStack struct {
    Head 		*Node
}


func Constructor() MinStack {
    return MinStack{}
}

func (this *MinStack) Push(val int)  {
	newNode := &Node{Value: val}
    if this.Head == nil {
		this.Head = newNode
		this.Head.MinmumValue = val
	} else {
		minVal := func() int {
			if val < this.Head.MinmumValue {
				return val
			}
			return this.Head.MinmumValue
		}

		newNode.MinmumValue = minVal()
		head := this.Head
		this.Head = newNode
		this.Head.Next = head
	}

}


func (this *MinStack) Pop()  {
    if this.Head == nil {
		return
	}
	this.Head = this.Head.Next
}


func (this *MinStack) Top() int {
 	return this.Head.Value
}


func (this *MinStack) GetMin() int {
	if this.Head != nil {
		return this.Head.MinmumValue
	}
	return 0
}

func (this *MinStack) print() {
	if this.Head == nil {
		fmt.Println("head", nil)
		return
	}
	curr := this.Head
	for curr != nil {
		fmt.Print(curr.Value, " ")
		curr = curr.Next
	}
	fmt.Println()
}


func aMainStack(){
	s := Constructor()
	s.Push(2147483646)
	s.print()
	
	s.Push(2147483646)
	s.print()
	
	s.Push(2147483647)
	s.print()
	

	s.Pop()
	s.print()
	

	s.Pop()
	s.print()
	

	s.Pop()
	s.print()
	

	s.Pop()
	s.print()
	



// 	s.Top()
// 	s.Pop() // 2
// 	// s.GetMin()
// 	// s.Pop()

// 	 s.print()
// 	//s.printMinStack()

// 	// s.GetMin()
// 	s.Pop() // 1
// // 
// 	s.print()
// 	//s.printMinStack()

// 	s.Push(2147483647) // 2
// 	// s.Top()
// 	// s.GetMin()
// 	s.Push(-2147483648) // 3
// 	s.print()
// 	//s.printMinStack()
// 	s.Top()
// 	// s.GetMin()
// 	s.Pop() //2
// 	// s.GetMin()
// 	s.print()
}