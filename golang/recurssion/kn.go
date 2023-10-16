package main

import "fmt"

func combine(n int, k int) [][]int{
	out := [][]int{}
	stack := []int{}
	backtrk(n, 1, k, stack, &out)

	fmt.Println(out)
	return out
}

func backtrk(n, start, k int, stack []int, out *[][]int){
	if len(stack) == k {
		*out = append(*out, stack)
		return
	}
	for i:=start; i< n+1; i++ {
		stack = append(stack, i)
		backtrk(n, i+1, k, stack, out)
		(stack) = (stack)[:len(stack)-1]
	}
}