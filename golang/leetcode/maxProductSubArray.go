package main

func maxProduct(nums []int) int {
	nLen := len(nums)
    lMSum := nums[0]
	rMSum := nums[ nLen - 1]
	lSum := 1
	rSum := 1

	for i := range nums {
		lSum *= nums[i]
		rSum *= nums[nLen - 1 - i]

		lMSum = maxN(lMSum, lSum)
		rMSum = maxN(rMSum, rSum)

		if lSum == 0 {
			lSum = 1
		}

		if rSum == 0{
			rSum = 1
		}
	}

	return maxN(lMSum, rMSum)
}

func maxN(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}