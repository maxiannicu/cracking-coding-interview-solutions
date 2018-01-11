package main

import "fmt"

type Stack struct {
	top *StackItem
}

type StackItem struct {
	next  *StackItem
	value int
}

func (this *Stack) Push(value int) {
	this.top = &StackItem{
		value: value,
		next:  this.top,
	}
}

func (this *Stack) Pop() int {
	val := this.top.value
	this.top = this.top.next

	return val
}

func (this *Stack) Peek() int {
	return this.top.value
}

func (this *Stack) IsEmpty() bool {
	return this.top == nil
}

// Runtime : O(n*n)
// Memory : O(n)
func sort(stack *Stack) {
	temp := Stack{}

	for {
		temp.Push(stack.Pop())
		for !stack.IsEmpty() && temp.Peek() < stack.Peek() {
			temp.Push(stack.Pop())
		}
		if stack.IsEmpty() {
			for !temp.IsEmpty() {
				stack.Push(temp.Pop())
			}
			break
		} else {
			divergedValue := stack.Pop()

			inserted := false
			for !temp.IsEmpty() {
				if divergedValue >= temp.Peek() && divergedValue < stack.Peek() {
					stack.Push(divergedValue)
					inserted = true
				} else {
					stack.Push(temp.Pop())
				}
			}

			if !inserted {
				stack.Push(divergedValue)
			}
		}
	}
}

func main() {
	stack := &Stack{}
	stack.Push(1)
	stack.Push(12)
	stack.Push(4)
	stack.Push(5)
	sort(stack)
	for !stack.IsEmpty() {
		fmt.Println(stack.Pop())
	}
}
