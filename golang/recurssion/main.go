package main

import "fmt"

func main() {
	//combine(4, 2)
	// ans := subsetsWithDup([]int{1,2,3,3})
	// fmt.Println(ans)

	// list := []int{1,2,3,4,5}
	// temp := make([]int, 3)
	// copy(temp, list)
	// fmt.Println(temp)

	numA1 := []int{15, -4, 56, 10, -23}
	numA2 := []int{14, -9, 56, 14, -23}
	ans := absSum2(numA1, numA2)
	fmt.Println(ans)

}
