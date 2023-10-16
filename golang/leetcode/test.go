package main

import "fmt"

type AFunc func(int, int)


func (a AFunc) Calculate(i int, j int) { 
	a(i, j)
}

func test(op string) AFunc {
	if op == "add"{
	return AFunc(func(i int, j int) {
		fmt.Println(i+j)
		})
	} 
	if op == "sub" {
		return AFunc(func(i int, j int) {
			fmt.Println(i-j)
		})
	}

	return AFunc(func(i int, j int) {
			fmt.Println(i*j)
	})
}