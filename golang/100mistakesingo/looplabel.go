package main

import "fmt"

func mainLoopLabel() {

loop:
	for i := 0; i < 10; i++ {
		switch i {
		default:
			fmt.Println(i)
		case 2:
			break loop
		}
	}

forloop:
	for i := 0; i < 3; i++ {
		fmt.Println("i", i)
		for j := 0; j < 3; j++ {
			fmt.Println("j", j)
			if j == 1 {
				break forloop
			}
		}
	}

}
