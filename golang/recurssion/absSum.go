package main

//numA1 := []int{15, -4, 56, 10, -23}
//numA2 := []int{14, -9, 56, 14, -23}

func absSum(numA1, numA2 []int) int {
	sum := 0
	bt(&sum, numA1, numA2)
	return sum
}

func bt(sum *int, numA1, numA2 []int)  {
	if len(numA1) == 0 {
		return
	}
	*sum += abs(numA1[0] - numA2[0])
	bt(sum, numA1[1:], numA2[1:])
}


func absSum2(numA1, numA2 []int) int {
	if len(numA1) == 0 {
		return 0
	} 
	return abs(numA1[0] - numA2[0]) + absSum2(numA1[1:], numA2[1:])
}



func abs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}