package main

import "fmt"

func firstCommonAscensor(root, p, q *TreeNode) *TreeNode {
	if root == nil || !isInSubTree(root, p) || !isInSubTree(root, q) {
		return nil
	}
	if p == q {
		return p
	}

	return firstCommonAscensorHelper(root, p, q)
}

func firstCommonAscensorHelper(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}

	pInLeft := isInSubTree(root.left, p)
	qInLeft := isInSubTree(root.left, q)

	if pInLeft != qInLeft {
		return root
	}

	if pInLeft {
		return firstCommonAscensorHelper(root.left, p, q)
	}

	return firstCommonAscensorHelper(root.right, p, q)
}

func isInSubTree(tree, node *TreeNode) bool {
	if tree == nil {
		return false
	}
	if tree == node {
		return true
	}

	return isInSubTree(tree.left, node) || isInSubTree(tree.right, node)
}

func main() {
	root := &TreeNode{
		value: 23,
		left: &TreeNode{
			value: 15,
			left: &TreeNode{
				value: 10,
				left: &TreeNode{
					value: 5,
				},
				right: &TreeNode{
					value: 13,
				},
			},
			right: &TreeNode{
				value: 20,
				left: &TreeNode{
					value: 17,
				},
				right: &TreeNode{
					value: 22,
				},
			},
		},
	}

	p := root.left.left
	q := root.left.right.left

	fmt.Println("First common ancestor between", p, "and", q, "is", firstCommonAscensor(root, p, q))
}
