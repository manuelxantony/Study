package main

func quicksort1(nums []int) {
	sort1(nums, 0, len(nums)-1)
}

func sort1(nums []int, low, high int) []int {
	if low < high{
		pivot := partition2(nums, low, high)
		sort1(nums, 0, pivot-1)
		sort1(nums, pivot+1, high)
	}
	return nums
}

func partition2(nums []int, low, high int) (int){
	pivot := high
	high--

	for  {
		for nums[low] < nums[pivot]{
			low ++
		}
		for nums[high] > nums[pivot]{
			high--
		}
		if low > high{
			break
		}
		nums[low], nums[high] = nums[high], nums[low]
		low ++
		high--
	}

	nums[pivot], nums[low] = nums[low], nums[pivot]
	return  low
}

// second way
func quickSort(nums []int) []int {
	sort(nums, 0, len(nums)-1)
	return nums
}

func sort(nums []int, lo, hi int){
	if lo < hi {
		pivot := partition(nums, lo, hi)
		sort(nums, lo, pivot-1)
		sort(nums, pivot+1, hi)
	}
}

func partition(nums []int, lo, hi int) int{
	pivot := nums[hi]
	for j:=lo; j < hi; j++{
		if nums[j] < pivot {
			nums[lo], nums[j] = nums[j], nums[lo]
			lo++
		}
	}
	nums[lo], nums[hi] = nums[hi], nums[lo]
	return lo
}