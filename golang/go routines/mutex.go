package main

import (
	"fmt"
	"sync"
	"time"
)


type Container struct {
	mu sync.Mutex
	balance int
}

func findBalance() {
	c := Container{
		balance: 0,
	}
	var wg sync.WaitGroup

	counter := func(){
		defer wg.Done()
		c.counter()
	}

	wg.Add(1)
	go counter()
	go counter()
	
	fmt.Println("waiting...")
	wg.Wait()
	
	// below lines will wait for go routines to finish
	fmt.Println("done")
	c.counter()
	fmt.Println(c.balance)
	
}

func (c *Container) counter() {
	c.mu.Lock()
	time.Sleep(time.Second * 2)
	for i:=0; i<10000000; i++{
		c.balance++
	}
	c.mu.Unlock()
}

