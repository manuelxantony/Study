package main

import "fmt"

func mainCall(){
	c := generator()
	receiver(c)
}

func receiver(c <- chan int){
	for o := range c {
		fmt.Println(o)
	}
}

func generator() <- chan int {
	c := make(chan int)
	go func(){
		for i:=0; i<10; i++ {
			c <- i
		}
		close(c)
	}()
	
	return c
}