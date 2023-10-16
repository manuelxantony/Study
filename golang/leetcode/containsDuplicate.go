package main


func containsDuplicate(nums []int) bool {
    numMap := make(map[int]int)
	for _, val := range nums {
		if _, ok := numMap[val]; ok{
			return true
		}
		numMap[val] = val
	}
	return false
}
