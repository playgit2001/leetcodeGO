package tree

func InorderTraversal(root *BinaryTreeNode) []int {
	var res []int
	inorderTraversalHelper(root, &res)
	return res
}

// 1递归法
func inorderTraversalHelper(root *BinaryTreeNode, res *[]int) {
	if root == nil {
		return
	}
	inorderTraversalHelper(root.Left, res)
	*res = append(*res, root.Val)
	inorderTraversalHelper(root.Right, res)
	return
}
