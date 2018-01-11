package main

import "fmt"

type Stack struct {
	top *StackItem
}

type StackItem struct {
	value int
	min int
	next *StackItem
}

func (this *Stack) Push(value int) {
	min := value
	if this.top != nil && this.top.min < min {
		min = this.top.min
	}
	this.top = &StackItem{
		value : value,
		next : this.top,
		min : min,
	}
}

func (this *Stack) Min() int {
	return this.top.min
}

func (this *Stack) Pop() int {
	value := this.top.value
	this.top = this.top.next

	return value
}

func main(){
	x := Stack{}
	x.Push(10)
	x.Push(4)
	x.Push(2)
	fmt.Println("Min is",x.Min())
	fmt.Println("Poping from stack...",x.Pop())
	fmt.Println("Min is",x.Min())
}
