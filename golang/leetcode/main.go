package main

import "fmt"

// type A struct{
// 	name string
// }

// type B struct {
// 	age int
// 	name string
// }

// func (a *A) Name() string{
// 	return a.name
// }

// func (b *B) Age() int {
// 	return b.age
// }

// func (b *B) Name() string{
// 	return b.name
// }

// type C struct{
// 	A
// 	B
// }

func main() {
	// c := &C{}
	// c.B.name = "name"
	// c.age = 20
	// fmt.Println(c.Age())
	// fmt.Println(c.B.Name())

	// two sum
	// nums := []int{2,7,11,15}
	// target := 9
	// ans := twoSum(nums, target)
	// fmt.Println(ans)

	// stock
	// prices := []int{7,1,5,3,6,4}
	// ans := maxProfit(prices)
	// fmt.Println(ans)

	// product of array expect self
	//nums := []int{-1,1,0,-3,3}
	// nums := []int{5,2, 3, 4}
	// ans := productExceptSelf(nums)
	// fmt.Println(ans)

	// max subarray
	// nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	// ans := maxSubArray(nums)
	// fmt.Println(ans)

	// contains duplicate
	// nums := []int{1,1,1,3,3,4,3,2,4,2}
	// ans := containsDuplicate(nums)
	// fmt.Println(ans)

	// search in rotated sorted array
	//nums := []int{6, 7, 1, 2, 3, 4, 5}
	// nums := []int{5,1,3}
	// target := 5
	// ans := search(nums, target)
	// fmt.Println(ans)

	// conatains most water
	// height := []int{1,1}
	// ans := maxArea(height)
	// fmt.Println(ans)

	// sliding window maximum
	// nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	// k := 3
	// ans := maxSlidingWindow(nums, k)
	// fmt.Println(ans)

	//test("add").Calculate(2,3)

	//executeThis()

	// queue
	// q := &Queue{}
	// q.Enqueue(1)
	// q.Enqueue(2)
	// q.print()

	// rune and string
	// s := string("dfdfd")
	// fmt.Println(len(s))
	// fmt.Println(utf8.RuneCountInString(s))

	// r := []rune(s)
	// fmt.Println(len(r))

	// pVal := []int{1,2,3,4}
	// val := 2

	// for _, pv := range pVal{
	// 	fmt.Println(pv == val)
	// }

	// s := "123"
	// fmt.Println(s)
	// modify(&s)
	// fmt.Println(s)

	// fmt.Println(IsValuePresent(3, 1, 2, 34))
	// fmt.Println(IsValuePresent("a", "v", "a", "d"))

	//fmt.Println( maxProduct([]int{-2, 3,  4}))

	//fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2, 3}, 6))
	//fmt.Println(search([]int{ 0, 1, 2, 3}, 3))

	// obj := Constructor()
	// p := obj.Insert(10)
	// fmt.Println(p)
	// p = obj.Insert(20)
	// fmt.Println(p)
	// p = obj.Insert(30)
	// p = obj.Insert(40)
	// p = obj.Insert(50)
	// p = obj.Insert(60)

	// fmt.Println(p)
	// fmt.Println(obj.GetRandom())

	// for _, v := range "abcd" {
	// 	fmt.Println(string(v))
	// }
	// canConstruct("aa", "ab")
	//testll()

	//mainSmall()

	// aMainStack()

	// arr := [][]int{{1,5,9},{10,11,13},{12,13,15}}
	// ans := kthSmallest(arr, 3)
	// fmt.Println(ans)

	// points := [][]int{{3, 3}, {5, -1}, {-2, 4}}
	// out := kClosest2(points, 2)
	// fmt.Println(out)
	// sli := make([]int, 0)
	// var sli2 []int
	// fmt.Println(cap(sli), cap(sli2))
	// if sli2 == nil {
	// 	fmt.Println(sli2, "sli2 is nill")
	// }
	// if sli == nil {
	// 	fmt.Println(sli, "sli is nill")
	// }

	// out := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	// fmt.Println(out)

	// m := Constructor1()
	// m.AddNum(2)
	// m.AddNum(1)
	// m.AddNum(5)
	// m.AddNum(7)
	// m.AddNum(3)
	// m.AddNum(4)
	// //m.AddNum(2)
	// fmt.Println(m.FindMedian())

	// grid := [][]byte{
	// 	{1, 0, 1, 1, 0, 1, 1},
	// }
	// fmt.Println(numIslands2(grid))

	grid2 := [][]string{{"1", "0", "0", "1", "1", "1", "0", "1", "1", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
		{"1", "0", "0", "1", "1", "0", "0", "1", "0", "0", "0", "1", "0", "1", "0", "1", "0", "0", "1", "0"},
		{"0", "0", "0", "1", "1", "1", "1", "0", "1", "0", "1", "1", "0", "0", "0", "0", "1", "0", "1", "0"},
		{"0", "0", "0", "1", "1", "0", "0", "1", "0", "0", "0", "1", "1", "1", "0", "0", "1", "0", "0", "1"},
		{"0", "0", "0", "0", "0", "0", "0", "1", "1", "1", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
		{"1", "0", "0", "0", "0", "1", "0", "1", "0", "1", "1", "0", "0", "0", "0", "0", "0", "1", "0", "1"},
		{"0", "0", "0", "1", "0", "0", "0", "1", "0", "1", "0", "1", "0", "1", "0", "1", "0", "1", "0", "1"},
		{"0", "0", "0", "1", "0", "1", "0", "0", "1", "1", "0", "1", "0", "1", "1", "0", "1", "1", "1", "0"},
		{"0", "0", "0", "0", "1", "0", "0", "1", "1", "0", "0", "0", "0", "1", "0", "0", "0", "1", "0", "1"},
		{"0", "0", "1", "0", "0", "1", "0", "0", "0", "0", "0", "1", "0", "0", "1", "0", "0", "0", "1", "0"},
		{"1", "0", "0", "1", "0", "0", "0", "0", "0", "0", "0", "1", "0", "0", "1", "0", "1", "0", "1", "0"},
		{"0", "1", "0", "0", "0", "1", "0", "1", "0", "1", "1", "0", "1", "1", "1", "0", "1", "1", "0", "0"},
		{"1", "1", "0", "1", "0", "0", "0", "0", "1", "0", "0", "0", "0", "0", "0", "1", "0", "0", "0", "1"},
		{"0", "1", "0", "0", "1", "1", "1", "0", "0", "0", "1", "1", "1", "1", "1", "0", "1", "0", "0", "0"},
		{"0", "0", "1", "1", "1", "0", "0", "0", "1", "1", "0", "0", "0", "1", "0", "1", "0", "0", "0", "0"},
		{"1", "0", "0", "1", "0", "1", "0", "0", "0", "0", "1", "0", "0", "0", "1", "0", "1", "0", "1", "1"},
		{"1", "0", "1", "0", "0", "0", "0", "0", "0", "1", "0", "0", "0", "1", "0", "1", "0", "0", "0", "0"},
		{"0", "1", "1", "0", "0", "0", "1", "1", "1", "0", "1", "0", "1", "0", "1", "1", "1", "1", "0", "0"},
		{"0", "1", "0", "0", "0", "0", "1", "1", "0", "0", "1", "0", "1", "0", "0", "1", "0", "0", "1", "1"},
		{"0", "0", "0", "0", "0", "0", "1", "1", "1", "1", "0", "1", "0", "0", "0", "1", "1", "0", "0", "0"}}

	fmt.Println(numIslands3(grid2))

}

// func IsValuePresent[T comparable] (value T, permittedValues ...T) bool {
// 	for _, pv := range permittedValues {
// 		if value == pv {
// 			return true
// 		}
// 	}
// 	return false
// }

// func modify(s *string) {
// 	*s = "oii"
// }
