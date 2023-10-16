package main

import (
	"fmt"
)

func main() {
	nums := Num{1, 2, 3}
	n := gimmeNum(nums) // interface call, nums []Num , where Num satisfies interface Number
	fmt.Println(n)
}
