package main

func subsetsWithDup(nums []int) [][]int {
	var queue [][]int
	queue = append(queue, []int{})
	counts := make(map[int]int)
	// for _, v := range nums{
	// 	counts[v]++
	// }
	for i:=0; i<len(nums); i++ {
		queueLen := len(queue)
		counts[nums[i]]++
		for j:=0; j<queueLen; j++{
			currrent := make([]int, len(queue[j]))
			copy(currrent, queue[j])
			if counts[nums[i]] == 1{
				currrent = append(currrent, nums[i])
				queue = append(queue, currrent)
			} else {
				
			}
		}
	}
	return queue
}