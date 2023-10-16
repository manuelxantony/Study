package main

import "fmt"

func panicrecover() {
	fmt.Println("going to panic")
	panic("panicked")

}

func mainpanicrecover() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover: ", r)
		}
	}()

	panicrecover()
	fmt.Println("I wont be executed")

}
