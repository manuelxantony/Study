package main

import "fmt"

func bubble(nums []int) []int{
	numLen := len(nums)
	for i:=0; i < numLen - 1; i++ {
		fmt.Println(nums)
		for j:=0; j < numLen - 1 - i; j++ {
			if nums[j] > nums[j+1]{
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}