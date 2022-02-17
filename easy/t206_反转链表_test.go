package easy

import (
	"testing"
)

/*
@Time : 2022/2/17 22:25
@Author : Firewine
@File : t206_反转链表_test.go
@Software: GoLand
@Description:
*/

func Test_reverseList(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l2 := l1
	l1.Next = &ListNode{Val: 2}
	l1 = l1.Next
	l1.Next = &ListNode{Val: 3}
	l1 = l1.Next
	l1.Next = &ListNode{Val: 4}
	l1 = l1.Next
	l1.Next = &ListNode{Val: 5}
	reverseList(l2)
}
