package main

import "fmt"

func firstMissingPositive(nums []int) int {
	small := 1
	lMap := make(map[int]int)

	for _, n := range nums {
		lMap[n] = n
	}

	for {
		if _, ok := lMap[small]; ok{
			small++
		} else {
			return small
		}
	}
}

func mainSmall(){
	l := []int{2, 1, -1, 0, 5, 4, 3}
	ans := firstMissingPositive(l)
	fmt.Println(ans)
}