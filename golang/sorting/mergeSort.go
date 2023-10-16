package main

import "fmt"


func mergeSort(nums []int) []int{
	ans := sortArray(nums)
	return ans

}
// 35, 33, 42, 10, 14, 19, 27, 44, 26, 31
// splitting here
func sortArray(nums []int) []int{
	numLen := len(nums)
	if numLen < 2 {
		return nums
	}
	fmt.Println(len(nums), numLen)

	first := sortArray(nums[:numLen/2])
	second := sortArray(nums[numLen/2:])	
	return merge(first, second)
}

func merge(first, second []int) []int{
	mergedSlice := []int{}
	l, r := 0, 0

	for l < len(first) && r < len(second){
		if first[l] < second[r]{
			mergedSlice = append(mergedSlice, first[l])
			l++
		} else {
			mergedSlice = append(mergedSlice, second[r])
			r++
		}
	}
	if l < len(first) {
		mergedSlice = append(mergedSlice, first[l:]...)
	
	}
	if r <len(second){
		mergedSlice = append(mergedSlice, second[r:]...)
	}

	return mergedSlice
}