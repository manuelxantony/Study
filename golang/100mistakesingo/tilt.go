package main

import (
	"fmt"
	"strconv"
)

type cInterface interface {
	~int
	String() string
}

type cInt int

func (i cInt) String() string {
	return strconv.Itoa(int(i))
}

func cmain() {
	var i cInt = 1
	fmt.Println(i.String())

}
