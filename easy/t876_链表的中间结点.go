package easy

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {

	if head.Next == nil {
		return head
	}

	left, right := head, head
	length := 1
	// 获取链表长度
	for right.Next != nil {
		length++
		right = right.Next
	}

	// 偶数，取中间第二个
	t := length/2 + 1
	for i := 1; i < t; i++ {
		left = left.Next
	}
	return left
}

type ListNode struct {
	Val  int
	Next *ListNode
}
