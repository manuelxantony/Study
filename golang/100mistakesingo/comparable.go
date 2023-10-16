package main

import (
	"fmt"
	"reflect"
)

func maincom() {
	s := []int{}
	s2 := []int{}
	fmt.Println(reflect.DeepEqual(s, s2))
}
