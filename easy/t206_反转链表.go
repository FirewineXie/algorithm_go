package easy

/*
@Time : 2022/2/17 22:21
@Author : Firewine
@File : t206_反转链表
@Software: GoLand
@Description:
*/

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev

}


