package main

import "fmt"

type Node struct {
	Value 		int
	MinimumValue int
}

type MinStack struct {
    Stack []Node
	PreviousMinimum int
}


func Constructor() MinStack {
    return MinStack{Stack: []Node{}}
}

func (this *MinStack) Push(val int)  {
	newNode := Node{Value: val}
	if len(this.Stack) == 0 {
		newNode.MinimumValue = val
		this.PreviousMinimum = val
	} else{
		if val < this.PreviousMinimum{
			this.PreviousMinimum = val
			newNode.MinimumValue = val
		} else{
			newNode.MinimumValue = this.PreviousMinimum
		}
	}
	this.Stack = append(this.Stack, newNode)

}


func (this *MinStack) Pop()  {
   	this.Stack = this.Stack[:len(this.Stack) - 1]
	len := len(this.Stack)
	if len >= 1{
   		this.PreviousMinimum = this.Stack[len - 1].MinimumValue
	}
}


func (this *MinStack) Top() int {
 return this.Stack[len(this.Stack) - 1].Value
}


func (this *MinStack) GetMin() int {
	return this.Stack[len(this.Stack) - 1].MinimumValue
}

func (this *MinStack) print() {
	for _, n := range this.Stack {
		fmt.Print("value: ", n.Value, "##","minValue: ", n.MinimumValue, "    ")
	}
	fmt.Println()
}


func aMainStack(){
	s := Constructor()
	s.Push(-124)
	s.Push(-164)
	s.print()
	fmt.Println(s.GetMin())
	fmt.Println(s.GetMin())
	fmt.Println(s.GetMin())
	s.Pop()
	s.print()
	fmt.Println(s.GetMin())
	s.Push(-24)
	s.Push(-100)
	s.print()
	fmt.Println(s.GetMin())
}