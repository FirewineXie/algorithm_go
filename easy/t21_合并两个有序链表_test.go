package easy

import (
	"testing"
)

/*
@Time : 2022/2/17 22:00
@Author : Firewine
@File : t21_合并两个有序链表_test.go
@Software: GoLand
@Description:
*/

func Test_mergeTwoLists(t *testing.T) {
	l1 := &ListNode2{Val: 1}
	l2 := l1
	l1.Next = &ListNode2{Val: 3}
	l1 = l1.Next
	l1.Next = &ListNode2{Val: 4}

	l3 := &ListNode2{Val: 1}
	l4 := l3
	l3.Next = &ListNode2{Val: 2}
	l3 = l3.Next
	l3.Next = &ListNode2{
		Val: 4,
	}
	mergeTwoLists(l2, l4)
}
