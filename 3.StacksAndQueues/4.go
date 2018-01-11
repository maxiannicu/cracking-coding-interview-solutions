package main

import "fmt"

type SimpleStack struct {
	top *SimpleStackItem
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
}

func (this *SimpleStack) Pop() interface{} {
	val := this.top.value
	this.top = this.top.next

	return val
}

func (this *SimpleStack) IsEmpty() bool {
	return this.top == nil
}

type MyQueue struct {
	stack SimpleStack
}

func (this *MyQueue) Enqueue(value interface{}) {
	temp := SimpleStack{}
	// move all items to temporary stack
	for !this.stack.IsEmpty() {
		temp.Push(this.stack.Pop())
	}

	// put new value in empty stack
	this.stack.Push(value)

	// put back values
	for !temp.IsEmpty() {
		this.stack.Push(temp.Pop())
	}
}

func (this *MyQueue) Dequeue() interface{} {
	return this.stack.Pop()
}

func main() {
	queue := MyQueue{}
	queue.Enqueue("first")
	queue.Enqueue("second")
	queue.Enqueue("third")
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
}
