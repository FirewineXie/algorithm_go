package easy

import (
	"fmt"
	"testing"
)

func Test_hasCycle(t *testing.T) {
	tt := &ListNode{Val: 1}

	fmt.Println(hasCycle(tt))

}
