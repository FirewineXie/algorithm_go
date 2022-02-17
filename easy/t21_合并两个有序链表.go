package easy

/*
@Time : 2022/2/17 21:47
@Author : Firewine
@File : t21_合并两个有序链表
@Software: GoLand
@Description:
*/

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func mergeTwoLists(list1 *ListNode2, list2 *ListNode2) *ListNode2 {

	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	topNode := &ListNode2{Val: 0}
	result := topNode
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			topNode.Next = list1
			list1 = list1.Next
			topNode = topNode.Next
			continue
		}
		if list1.Val > list2.Val {
			topNode.Next = list2
			list2 = list2.Next
			topNode = topNode.Next
			continue
		}

	}
	if list1 != nil {
		topNode.Next = list1
	}
	if list2 != nil {
		topNode.Next = list2
	}

	return result.Next
}
