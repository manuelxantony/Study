package main

func productExceptSelf(nums []int) []int {
    ans := []int{}
    prefix := 1

    for _, num := range nums {
        ans = append(ans, prefix)
        prefix *= num
    }

	postfix := 1
	lenNums := len(nums)

	for i := range nums {
		index := lenNums - i - 1
		ans[index] *= postfix
		postfix *= nums[index]
	}

	return ans
}