package main

import "fmt"

type InterfaceA interface {
	Len() int
}

type InterfaceB interface {
	Size() int
}

type A struct {
	len int
}

type B struct {
	A
}

type C struct {
	A
	B
}

func (a A) Len() int {
	return a.len
}

// func (b B) Len() int {
// 	return b.len
// }

func print(i InterfaceA) {
	fmt.Println("len:", i.Len(), i)
}

func main1() {
	a := A{5}
	b := B{a}
	c := C{a, b}
	print(b)
	print(c)
}
