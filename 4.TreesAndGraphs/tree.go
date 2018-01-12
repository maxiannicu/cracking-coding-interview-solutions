package main

import "fmt"

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func getChildDescriptor(child *TreeNode) string {
	if child == nil {
		return "-"
	}
	return "+"
}

func (this *TreeNode) String() string {
	return fmt.Sprintf("%d[%v,%v]", this.value, getChildDescriptor(this.left), getChildDescriptor(this.right))
}

func (this *TreeNode) PrintInOrder() {
	if this.left != nil {
		this.left.PrintInOrder()
	}
	fmt.Printf("%v ", this)
	if this.right != nil {
		this.right.PrintInOrder()
	}
}

func (this *TreeNode) PrintPreOrder() {
	fmt.Printf("%v ", this)
	if this.left != nil {
		this.left.PrintPreOrder()
	}
	if this.right != nil {
		this.right.PrintPreOrder()
	}
}

func (this *TreeNode) PrintPostOrder() {
	if this.left != nil {
		this.left.PrintPostOrder()
	}
	if this.right != nil {
		this.right.PrintPostOrder()
	}
	fmt.Printf("%v ", this)
}
