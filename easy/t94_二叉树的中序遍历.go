package easy

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return dfsInorderTraversal(root)
}

func dfsInorderTraversal(node *TreeNode) (result []int) {

	if node.Left == nil && node.Right == nil {
		result = append(result, node.Val)
		return result

	}

	if node.Left != nil {
		result = append(result, dfsInorderTraversal(node.Left)...)

	}
	result = append(result, node.Val)
	if node.Right != nil {
		result = append(result, dfsInorderTraversal(node.Right)...)
	}
	return result
}
