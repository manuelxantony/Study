package main

import "fmt"

func usingAny() {
	var i any
	i = 1
	i = "any"
	i = 3.0
	i = struct {
		s string
	}{
		s: "ohhh",
	}
	fmt.Println(i)
	// for compiling
	_ = i

}
