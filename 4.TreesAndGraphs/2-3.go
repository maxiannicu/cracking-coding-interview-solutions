package main

import "fmt"

func createTree(arr []int) *TreeNode {
	length := len(arr)
	middle := length / 2
	if length == 0 {
		return nil
	}
	if length == 1 {
		return &TreeNode{
			value: arr[0],
		}
	}
	if length%2 == 0 {
		middle--
	}

	return &TreeNode{
		left:  createTree(arr[0:middle]),
		right: createTree(arr[middle+1 : length]),
		value: arr[middle],
	}
}

func scanDepth(root *TreeNode, level int, result map[int][]*TreeNode) {
	if root == nil {
		return
	}
	if _, ok := result[level]; !ok {
		result[level] = make([]*TreeNode, 0)
	}

	result[level] = append(result[level], root)

	scanDepth(root.left, level+1, result)
	scanDepth(root.right, level+1, result)
}

func getDepths(root *TreeNode) map[int][]*TreeNode {
	result := map[int][]*TreeNode{}
	scanDepth(root, 0, result)
	return result
}

func main() {
	arr := []int{5, 10, 13, 15, 17, 20, 22, 23, 24, 27, 29, 50, 55, 75, 80}
	root := createTree(arr)

	fmt.Println("Tree inorder : ")
	root.PrintInOrder()
	fmt.Println()

	fmt.Println("Tree preorder : ")
	root.PrintPreOrder()
	fmt.Println()

	fmt.Println("Tree postorder : ")
	root.PrintPostOrder()
	fmt.Println()

	fmt.Println("Scanning depth")
	result := getDepths(root)
	for key, value := range result {
		fmt.Printf("%d ", key)
		for _, el := range value {
			fmt.Printf("%v", el)
		}
		fmt.Println()
	}
}
