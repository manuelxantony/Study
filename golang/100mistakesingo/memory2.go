package main

import (
	"fmt"
	"runtime"
)

func memorymain2() {
	foos := [][]byte{}
	printMem2() // 106 KB

	for i := 0; i < 1000; i++ {
		foos = append(foos, make([]byte, 1024*1024))
	}
	printMem() // 1024109 KB

	two := keeptwoelementsonly2(foos)
	runtime.GC()
	printMem2()            // 1024110 KB
	runtime.KeepAlive(two) // makes sure that two is not cleared until here
}

func printMem2() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)
}

func keeptwoelementsonly2(foos [][]byte) [][]byte {
	return foos[:2]
}
