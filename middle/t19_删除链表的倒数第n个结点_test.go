package middle

import (
	"fmt"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	head := ListNode{}
	right := &head
	n := 2
	for i := 1; i <= n; i++ {
		right.Val = i
		if i != n {
			right.Next = &ListNode{}
		}

		right = right.Next
	}
	fmt.Println(removeNthFromEnd(&head, 2))
}
