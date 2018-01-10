package main

import "fmt"
import "math/rand"


type OneArrayStacks struct {
	values []int
	references []int
	stackTops []int
	top int
}

// Runtime : O(n)
// Memory : O(n)
func NewOneArrayStacks(stacks,capacity int) *OneArrayStacks {
	capacity++
	values := make([]int, capacity)
	references := make([]int, capacity)
	
	references[0] = -1
	references[capacity-1] = -1
	for i := 1; i < capacity - 1; i++ {
		references[i] = i+1
	}

	stackTops := make([]int, stacks)
	for i:= 0; i < stacks; i++ {
		stackTops[i] = -1
	}

	return &OneArrayStacks{
		values : values,
		references : references,
		stackTops : stackTops,
		top : 1,
	}
}

// Runtime : O(1)
// Memory : O(1)
func (this *OneArrayStacks) Push(stack,value int){
	newTop := this.references[this.top]
	
	this.references[this.top] = this.stackTops[stack]
	this.values[this.top] = value
	this.stackTops[stack] = this.top

	this.top = newTop
}

// Runtime : O(1)
// Memory : O(1)
func (this *OneArrayStacks) Pop(stack int) int {
	top := this.stackTops[stack]
	this.stackTops[stack] = this.references[top]
	val := this.values[top]
	
	this.references[top] = this.top
	this.top = top

	return val
}


// Runtime : O(1)
// Memory : O(1)
func (this *OneArrayStacks) IsFull() bool {
	return this.top == -1
}

// Runtime : O(1)
// Memory : O(1)
func (this *OneArrayStacks) IsEmpty(stack int) bool {
	return this.stackTops[stack] == -1
}

func main(){
	stacks := 3
	obj := NewOneArrayStacks(stacks,10)

	for ; !obj.IsFull() ; {
		s := rand.Int() % stacks
		v := rand.Int() % 1012
		obj.Push(s,v)
		fmt.Println("Inserted value",v,"in stack",s)
	}

	fmt.Println("Clearing stacks...")
	for i := 0 ; i < stacks ; i++ {
		fmt.Println("Clearing stack ",i)
		for ; !obj.IsEmpty(i) ; {
			fmt.Println(obj.Pop(i))
		}
	}
}
