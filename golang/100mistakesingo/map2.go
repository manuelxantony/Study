package main

import "fmt"

func mainmap() {
	m := make(map[string]int)
	m["one"] = 1
	m["tw0"] = 2
	m["three"] = 3
	m["four"] = 4

	fmt.Println(m["one"])

}
