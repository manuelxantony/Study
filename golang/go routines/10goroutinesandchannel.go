package main

import (
	"fmt"
)


func mainHere(){
	sy := &SyncProperties{ch:make(chan any)}
	for i:=0; i<10; i++ {
		sy.wg.Add(1)
		go sy.gimmeHello()
		
	}
	// problem of calling sy.wg.Wait() here is that,
	// go routines are been called from for loop
	// and even before the sy.wg.Done() is been executed sy.wg.Wait() will execute
	// but there are no routines to wait
	// sy.wg.Wait()

	// close the channle after wait
	// this wait will wait for all go routines to finish then
	// close will close the channel
	go func() {
		sy.wg.Wait()
		close(sy.ch)
	}()

	for out := range sy.ch {
		fmt.Println(out)
	}
	
}

func (sy *SyncProperties) gimmeHello(){
	defer sy.wg.Done()
	sy.ch <- "hello"
}