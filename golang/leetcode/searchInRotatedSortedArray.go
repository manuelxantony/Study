package main

// func search(nums []int, target int) int {

// 	 low, high := 0, len(nums) -1

// 	 for low <= high{
// 		mid := (low + high) / 2

// 		if target == nums[mid]{
// 			return mid
// 		}

// 		if nums[low]<= nums[mid] {
// 			if target >= nums[low] && target < nums[mid]{
// 				high = mid - 1
// 			} else{
// 				low = mid + 1
// 			}
// 		} else {
// 			if target > nums[mid] && target <= nums[high]{
// 				low = mid +  1
// 			} else{
// 				high = mid - 1
// 			}
// 		}
// 	}

// 	return -1
// }

// different than the question
// here trying to find the exact position if not rotated
// func searchInOrder(nums []int, target int) int {
// 	l, r := 0, len(nums) - 1
// 	tPos, smallPos := -1, -1

// 	// if rotated
// 	if nums[l] > nums[r] {
// 		for i := range nums {
// 			if nums[i] == target {
// 				tPos = i
// 			}
// 			if nums[l] > nums[i] && smallPos == -1{
// 				smallPos = i
// 			}
// 			if tPos != -1 && smallPos != -1 {
// 				if smallPos <= tPos{
// 					return tPos - smallPos
// 				} else {
// 					return len(nums) - smallPos + tPos
// 				}
// 			}
// 		}
// 	} else {
// 		// not rotated
// 		for i := range nums {
// 			if nums[i] == target {
// 				return i
// 			}
// 		}
// 	}
// 	return -1
// }

// func search(nums []int, target int) int {
// 	l, r := 0, len(nums) - 1

// }