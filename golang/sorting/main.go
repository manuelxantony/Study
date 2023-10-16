package main

import "fmt"

func main(){
	// arr := []int{6,5,4,3,2,1}
	arr := []int{35, 33, 42, 10, 14, 19, 27, 44, 26, 31}
	//quicksort1(arr)
	ans := mergeSort(arr)
	fmt.Println(ans)
}