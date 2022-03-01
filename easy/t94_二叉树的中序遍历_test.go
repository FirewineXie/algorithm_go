package easy

import (
	"fmt"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	root := &TreeNode{}

	root.Val = 1
	root.Right = &TreeNode{Left: &TreeNode{Val: 3}, Val: 2}

	traversal := inorderTraversal(root)
	fmt.Println(traversal)
}
