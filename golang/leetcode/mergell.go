package main

import "fmt"

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	mL := head

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			mL.Next = list1
			list1 = list1.Next
		} else {
			mL.Next = list2
			list2 = list2.Next
		}
		mL = mL.Next
	}

	if list1 == nil {
		mL.Next = list2
	}
	if list2 == nil {
		mL.Next = list1
	}

	return head.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	mergedList := lists[0]

	if len(lists) == 1 {
		return mergedList
	}

	for _, list := range lists[1:] {
		if list != nil {

			mergedList = mergeTwoLists(mergedList, list)
		}
	}

	return mergedList
}

func testll() {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 3}
	l1.Next = l2
	l2.Next = l3

	l21 := &ListNode{Val: 4}
	l22 := &ListNode{Val: 5}
	l23 := &ListNode{Val: 6}
	l21.Next = l22
	l22.Next = l23

	//l31 := &ListNode{}

	list := []*ListNode{}
	list = append(list, l1)
	list = append(list, nil)
	list = append(list, l21)
	ans := mergeKLists(list)

	for ans != nil {
		fmt.Println(ans.Val)
		ans = ans.Next
	}

}
