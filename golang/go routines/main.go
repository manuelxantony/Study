package main

// https://www.golangprograms.com/goroutines-and-channels-example.html

import (
	"fmt"
	"sync"
)

type SyncProperties struct {
	wg sync.WaitGroup
	ch chan any
}

// func f(s string) {
// 	for i:=0; i<3; i++ {
// 		fmt.Printf("%s : %d\n", s, i)
// 	}
// }

func main() {
	// go f("routine world")
	
	// f("hello")
	
	// go func(t string){
	// 	fmt.Println(t)
	// }("going")

	// time.Sleep(time.Second)
	// fmt.Println("done")

	// interesting_problem()
	// interesting_solution()

	//channel()

	// findBalance()
	// mainHere()
	// mainCall()
	// mainFib(20)

}

// func interesting_problem() {
// 	var wg sync.WaitGroup
// 	for _, value := range []string{"apple", "orange", "mango"} {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			fmt.Println(value)
// 		}()
// 	}
// 	wg.Wait()
// }

// func interesting_solution() {
// 	var wg sync.WaitGroup
// 	for _, salutation := range []string{"apple", "orange", "mango"} {
// 		wg.Add(1)
// 		go func(salutation string){
// 			defer wg.Done()
// 			fmt.Println(salutation)
// 		}(salutation)
// 	}
// 	wg.Wait()
// }

// channel

func channel() {
	var wg sync.WaitGroup

	msg := make(chan string)
	var greeting string
	wg.Add(1)
	go greet(msg)
	defer wg.Done()

	greeting = <- msg
	fmt.Println("Greeting received")

	fmt.Println(greeting)

	if _, ok := <- msg; ok {
		fmt.Println("Chanel is open")
	} else {
		fmt.Println("Channel is closed")
	}

	
}

func greet(ch chan string) {
	fmt.Println("Waiting to send greeting")
	ch <- "hello, blue"
	close(ch)
	fmt.Println("Greeting sent")
}