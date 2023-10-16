package main

func maxArea(height []int) int {
 	l, r := 0, len(height) - 1
 	maxArea := 0
	area := 0
	
 	for l < r {
		if height[l] >= height[r] {
				area = ((r-l) * (min(height[l], height[r])))
			r -= 1
		} else{
				area = ((r-l) * (min(height[l], height[r])))
			l += 1
		}
			maxArea = max(maxArea, area)

 	}
	
	return maxArea
}

func min(n1 int, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func max(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}