package main

import (
	"fmt"
	"runtime"
)

type Foo struct {
	v []byte
}

func memorymain() {
	foos := make([]Foo, 1_000)
	printMem() // 106 KB

	for i := 0; i < len(foos); i++ {
		foos[i] = Foo{
			v: make([]byte, 1024*1024),
		}
	}
	printMem() // 1024109 KB

	two := keeptwoelementsonly(foos)
	runtime.GC()
	printMem()             // 1024110 KB
	runtime.KeepAlive(two) // makes sure that two is not cleared until here

	// clearing
	runtime.GC()
	// fixed
	two = keeptwoelementsonlyFIXED(foos)
	runtime.GC()
	printMem()
	runtime.KeepAlive(two) // 2162

}

func printMem() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)
}

func keeptwoelementsonly(foos []Foo) []Foo {
	return foos[:2]
}

func keeptwoelementsonlyFIXED(foos []Foo) []Foo {
	newFoos := make([]Foo, 2)
	copy(newFoos, foos)
	return newFoos
}
