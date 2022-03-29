package middle

import (
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	node := ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	node1 := ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	addTwoNumbers(&node, &node1)
}
