package main

//1, 3, -1, -3, 5, 3, 6, 7
func maxSlidingWindow(nums []int, k int) []int {
    out := []int{}
    q := []int{}

    for i, n := range nums {

        for len(q) > 0 && nums[q[0]] < n {
            q = q[1:]
        }

        q = append(q, i)
        
        if i < k - 1 {
            continue
        }

        out = append(out, nums[q[0]])

        if q[0] == i - k + 1 {
            q = q[1:]
        }


    }

    

    return out
}