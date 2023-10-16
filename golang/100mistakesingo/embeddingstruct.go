package main

import (
	"fmt"
)

type Base struct {
	b int
}

type Container struct {
	Base
	b string
}

func (b Base) Describe() {
	fmt.Println(b.b)
}

// func (b Container) Describe() {
// 	fmt.Println(b.b)
// }

func embStruct() {
	c := Container{}
	c.b = "abc"
	c.Base.b = 234
	fmt.Println(c)

	c2 := Container{Base: Base{123}, b: "abc"}
	//if receiver Describe for Container was there then,
	//c2.Describe() // out -> "abc"
	// because child have more presidence

	c2.Describe() // out -> 123

}
