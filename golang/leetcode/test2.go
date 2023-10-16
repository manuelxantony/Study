package main

import "fmt"

func add(i int, j int) int {
	return i + j
}

type Addition func(int, int) int

func executeThis() {
	aArr := []Addition{}
	for range [10]int{} {
		aArr =  append(aArr, add)
	}

	for i, a := range aArr {
		h := a(1, i)
		fmt.Println(h)
	}

}