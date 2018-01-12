package main

import "math"
import "fmt"

func isBalancedSubTree(root *TreeNode, depth int) (int, bool) {
	if root == nil {
		return depth - 1, true
	}
	leftDepth, leftBalanced := isBalancedSubTree(root.left, depth+1)
	rightDepth, rightBalanced := isBalancedSubTree(root.right, depth+1)
	isBalanced := leftBalanced && rightBalanced && (math.Abs(float64(leftDepth)-float64(rightDepth)) < 2)
	subtreeDepth := leftDepth
	if rightDepth > depth {
		subtreeDepth = rightDepth
	}

	return subtreeDepth, isBalanced
}

func isBalanced(root *TreeNode) bool {
	_, ok := isBalancedSubTree(root, 0)

	return ok
}

func main() {
	root := &TreeNode{
		value: 50,
		left: &TreeNode{
			value: 40,
			left: &TreeNode{
				value: 30,
			},
		},
	}

	fmt.Println(isBalanced(root))
}
