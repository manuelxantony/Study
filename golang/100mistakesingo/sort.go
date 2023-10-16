package main

import (
	"fmt"
	"sort"
)

func mainSort() {
	arr := []int{1, 2, 3, 4}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	fmt.Println(arr)
}
