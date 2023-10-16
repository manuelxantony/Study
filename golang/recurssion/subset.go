package main

// func subsets2(nums []int) [][]int {
//     var queue [][]int
//     queue = append(queue, []int{})

//     for i := 0; i < len(nums); i++ {
//         length := len(queue)

//         for j := 0; j < length; j++ {
//             item := queue[j]

//             tmp := make([]int, len(item))
//             copy(tmp, item)
//             tmp = append(tmp, nums[i])

//             queue = append(queue, tmp)
//         }
//     }

//     return queue
// }

func subsets2(nums []int) [][]int {
	var queue [][]int
	queue = append(queue, []int{})

	for i:=0; i<len(nums); i++ {
		l := len(queue)
		for j:=0; j<l; j++{
			item := make([]int, len(queue[j]))
			copy(item, queue[j])
			item = append(item, nums[i])
			queue = append(queue, item)
		}
	}
	return queue
}


func subsets(nums []int) [][]int {
	var queue [][]int
	queue = append(queue, []int{})

	for i:=0; i<len(nums); i++{
		lenQueue := len(queue)
		for j:=0; j<lenQueue; j++{ // len(queue) -> infinite loop
			temp := make([]int, len(queue[j]))
			copy(temp, queue[j]) 
			temp = append(temp, nums[i])
			queue = append(queue, temp)
		}
	}
	return queue
}

// func subsetsBT(nums []int) [][]int{
// 	var queue [][]int
// 	backtracking(nums, &queue, 0, []int{})
// 	return queue
// }

// func backtracking(nums []int ,queue *[][]int, start int, current []int) {
// 	temp := make([]int, len(current))
// 	copy(temp, current)
// 	*queue = append(*queue, temp)

// 	for i:=start; i<len(nums); i++{
// 		current = append(current, nums[i])
// 		backtracking(nums, queue, i+1, current)
// 		current = current[:len(current)-1]
// 	}
// }

