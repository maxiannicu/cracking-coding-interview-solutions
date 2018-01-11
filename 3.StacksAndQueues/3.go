package main

import "fmt"

type Stack interface {
	Push(value interface{})
	Pop() interface{}
	Peek() interface{}
	Capacity() int
	Items() int
}

type SimpleStack struct {
	top      *SimpleStackItem
	capacity int
	items    int
}

type SimpleStackItem struct {
	next  *SimpleStackItem
	value interface{}
}

func (this *SimpleStack) Push(value interface{}) {
	this.top = &SimpleStackItem{
		value: value,
		next:  this.top,
	}
	this.items++
}

func (this *SimpleStack) Pop() interface{} {
	val := this.top.value
	this.top = this.top.next
	this.items--

	return val
}

func (this *SimpleStack) Peek() interface{} {
	return this.top.value
}

func (this *SimpleStack) Capacity() int {
	return this.capacity
}

func (this *SimpleStack) Items() int {
	return this.items
}

type StackOfStacks struct {
	stack SimpleStack
}

func (this *StackOfStacks) Push(value interface{}) {
	if this.stack.Items() == 0 || this.stack.Peek().(Stack).Items() == this.stack.Peek().(Stack).Capacity() {
		fmt.Println("creating new stack")
		this.stack.Push(&SimpleStack{
			capacity: 5,
		})
	}
	this.stack.Peek().(Stack).Push(value)
}

func (this *StackOfStacks) Pop() interface{} {
	if this.stack.Peek().(Stack).Items() == 0 {
		this.stack.Pop()
		fmt.Println("removing empty stack")
	}

	return this.stack.Peek().(Stack).Pop()
}

func main() {
	stacks := StackOfStacks{}
	stacks.Push(1)
	stacks.Push(5)
	stacks.Push(2)
	stacks.Push(2)
	stacks.Push(3)
	stacks.Push(5) // should create new stack as previous is full
	stacks.Push(77)
	stacks.Pop()
	stacks.Pop() // should remove last stack, as it is empty
	stacks.Pop()
}
