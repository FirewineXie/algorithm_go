package easy

import (
	"fmt"
	"testing"
)

func Test_middleNode(t *testing.T) {
	head := ListNode{}
	right := &head
	for i := 1; i <= 6; i++ {
		right.Val = i
		if i != 6 {
			right.Next = &ListNode{}
		}

		right = right.Next
	}
	fmt.Println(middleNode(&head))

}
