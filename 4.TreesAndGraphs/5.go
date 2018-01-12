package main

import "fmt"

func isBinaryTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	result := true

	if root.left != nil {
		result = result && isBinaryTree(root.left) && root.value >= root.left.value
	}
	if root.right != nil {
		result = result && isBinaryTree(root.right) && root.value < root.right.value
	}

	return result
}

func main() {
	root := &TreeNode{
		value: 3,
		left: &TreeNode{
			value: 1,
		},
	}

	fmt.Println(isBinaryTree(root))
}
