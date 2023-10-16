package main

func kthSmallest(matrix [][]int, k int) int {
	matrixLen := len(matrix)
	l, r := matrix[0][0], matrix[matrixLen-1][matrixLen-1]
	
	for l < r{
		mid := (r+l)/2
		count := lessCount(matrix, mid)
		if count < k {
			l = mid+1
		} else if count >= k{
			r = mid 
		} 
	}

	//[1,  5,  9]
	//[10, 11, 13]
	//[12, 13, 15]
	// after mid find how many numbers are smaller or equal than mid

	return l
	
}

func lessCount(matrix [][]int ,num int) int {
	count := 0
	for _, mat := range matrix {
		for _, val := range mat {
			if val <= num {
				count++
			}
		}
	}
	return count
}

// quick sort
// func sort(first, second []int) []int {
// 	// treat them like unsorted array
// 	// not the best approach
// 	arr := []int{}
// 	arr = append(arr, first...)
// 	arr = append(arr, second...)
// 	quickSort(arr, 0, len(arr)-1)
// 	return arr
// }

// func quickSort(arr []int, low, high int) {
// 	if low < high{
// 		pivot := partition(arr, low, high)
// 		quickSort(arr, low, pivot-1)
// 		quickSort(arr, pivot+1, high)
// 	}
// }

// func partition(arr []int, low, high int) int{
// 	pivot := high
// 	for i:=low; i< high; i++ {
// 		if arr[i] < arr[pivot]{
// 			arr[low], arr[i] = arr[i], arr[low]
// 			low ++
// 		}
// 	}
// 	arr[low], arr[high] = arr[high], arr[low]
// 	return low
// }