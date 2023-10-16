package main

// n^2
func twoSumN2(nums []int, target int) []int {
	for i :=0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i!= j &&target == nums[i] + nums[j]{
				return []int{i, j}
			}
		}
	}

	return []int{}
}

// n 
func twoSum(nums []int, target int) []int{
	iMap := make(map[int]int)
	for i, num := range nums {
		lookFor := target - num
		if val, isFound := iMap[lookFor]; isFound {
			return []int{i, val}

		}
		iMap[num] = i
	}

	return []int{}
}