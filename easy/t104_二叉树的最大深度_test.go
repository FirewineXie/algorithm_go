package easy

import (
	"fmt"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}

	a := maxDepth(root)
	fmt.Println(a)

}
