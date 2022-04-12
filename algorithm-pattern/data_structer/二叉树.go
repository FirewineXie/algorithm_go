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

/*
遍历相当于 ： 每个节点，都有 左，中 右 ，，
前序遍历：当到达这个节点，如果是前序遍历，顺序 根左右
中序遍历：当到达这个节点，如果是中序遍历，顺序 左根右
后序遍历：当到达这个节点，如果是后序遍历，顺序 左右根

*/

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

// 中序非递归
// 利用stack 进行回溯
func inorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left // 一直想左
		}

		// 弹出
		val := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, val.Val)
		root = val.Right
	}

	return result
}

// 后序非递归
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var lastVisit *TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		// 判断当前结点是否有右节点
		val := stack[len(stack)-1]
		if val.Right == nil || val.Right == lastVisit {
			stack = stack[:len(stack)-1] //pop
			result = append(result, val.Val)
			// 标记当前节点已经弹出过
			lastVisit = val
		} else {
			root = val.Right
		}
	}

	return result
}

//分治法

func preorderTraversalV2(root *TreeNode) []int {
	result := divideAndConquer(root)
	return result
}

func divideAndConquer(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	// divide : 是将大的问题，平均切割每个小问题，每个小问题合起来是大问题的解
	left := divideAndConquer(root.Left)
	right := divideAndConquer(root.Right)

	// 合并问题的解  ,根左右
	result = append(result, root.Val)
	result = append(result, left...)
	result = append(result, right...)

	return nil
}

// 层次遍历 BFS 广度优先算法

func levelOrder(root *TreeNode) [][]int {
	// 通过上一层的长度确定下一层的元素
	result := make([][]int, 0)

	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		list := make([]int, 0)
		// 记录当前层有多少元素（遍历当前层，再添加下一层）
		l := len(queue)
		for i := 0; i < l; i++ {
			// 出队列
			level := queue[0]
			queue = queue[1:]
			list = append(list, level.Val)
			if level.Left != nil {
				queue = append(queue, level)
			}
			if level.Right != nil {
				queue = append(queue, level)
			}
		}
		result = append(result, list)
	}
	return result
}
