package main

import "fmt"

type Number interface {
	gimmeNum() int
}

func gimmeNum(n Number) int {
	fmt.Println("called into interface")
	return n.gimmeNum()
}
