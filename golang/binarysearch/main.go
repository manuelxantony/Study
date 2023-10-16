package main

import "fmt"


func main() {
	arr := []int{1,2, 9, 20, 31, 45, 63, 70, 100}
	target := 70
	if ok, pos := binarySearchReccursion(arr, target, 0); ok{
		fmt.Printf("target %d in position: %d\n", target, pos)
	} else{
		fmt.Println("Not found in array")
	}
	
	
}

func binarySearch(arr []int, target int) (bool, int){
	arrLen := len(arr)
	l, r := 0, arrLen-1

	for l <= r{
		mid := (r+l)/2

		if arr[mid] == target{
			return true, mid
		}
		if target < arr[mid]{
			r = mid-1

		}
		if target > arr[mid]{
			l = mid+1
		}
	}
	return false, -1
}

func binarySearchReccursion(arr []int, target int, pos int) (bool, int){
	l ,r := 0, len(arr)-1
	mid := (r+l)/2
	found, index := false, -1

	if l <= r{
		if arr[mid] == target{
			return true, mid + pos
		}
		if target < arr[mid]{
			found, index = binarySearchReccursion(arr[:mid], target, 0)
		}
		if target > arr[mid]{
			found, index = binarySearchReccursion(arr[mid+1:], target, pos + mid+1)
		}
	}
	return found, index
}