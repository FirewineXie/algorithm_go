package data_structer

import (
	"fmt"
)

/*
知识点：
前序遍历： 先访问根节点，再前序遍历左子树，再前序遍历右子树
中序遍历： 先中序遍历左子树，再访问根节点，再中序遍历右子树
后序遍历： 先后序遍历左子树，再后序遍历右子树，再访问根节点
层序遍历： 先一层根节点访问，再每一层挨个访问节点

注意点：
1. 访问根节点时间，来决定是什么遍历
2. 左子树都是优先于右子树

*/

//TreeNode 结构体
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序递归
// eg: 中序递归 & 后序遍历，只需要改变 输出的code 所在顺序
func preorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	// 先访问根节点
	fmt.Println(root)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
}

// 前序非递归
func preorderTraversalV1(root *TreeNode) {
	if root == nil {
		return
	}
	var result []int
	var stack []*TreeNode

	for root != nil && len(stack) != 0 {
		for root != nil {
			// 前序遍历，先保存结果
			result = append(result, root.Val)
			stack = append(stack, root.Left)
			root = root.Left
		}
		// pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
}
