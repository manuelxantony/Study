package main

type Num []int

func (n Num) gimmeNum() int {
	if len(n) == 0{
		return -1
	}
	return n[0]
}