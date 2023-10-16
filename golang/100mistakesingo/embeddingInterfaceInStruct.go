package main

import "fmt"

type Interface interface {
	FirstSecond(string, string)
}

type Struct struct{}

func (t Struct) FirstSecond(s1, s2 string) {
	fmt.Println(s1, s2)
}

type Reverse struct {
	Interface
}

func (c Reverse) FirstSecond(s1, s2 string) {
	// here c.Interface is Struct
	c.Interface.FirstSecond(s2, s1)
}

func PrintName(I Interface, s1, s2 string) {
	I.FirstSecond(s1, s2)
}

func embdInterfaceInStruct() {
	s := Struct{}
	r := Reverse{s}
	PrintName(s, "first", "second")
	PrintName(r, "first", "second")
}
