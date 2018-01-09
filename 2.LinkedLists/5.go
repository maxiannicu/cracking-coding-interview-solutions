package main

import "fmt"
import "github.com/golang-collections/collections/stack"


// Runtime : O(a) || O(b) . Could not be determined as we don't have relationship between A and B
// Memory : O(a) || O(b) . Could not be determined as we don't have relationship between A and B
func sum(a,b *SingleLinkedListNode) *SingleLinkedListNode {
	var lastResultDigit *SingleLinkedListNode
	result := &SingleLinkedListNode{
		value : 0,
	}
	carry := 0

	for ; a != nil || b != nil ; {
		newDigit := carry
		if a != nil {
			newDigit += a.value
			a = a.next
		}
		if b != nil {
			newDigit += b.value
			b = b.next
		}
		carry = newDigit / 10
		newDigit = newDigit % 10
		if lastResultDigit == nil {
			result.value = newDigit
			lastResultDigit = result
		} else {
			lastResultDigit.next = &SingleLinkedListNode{
				value : newDigit,			
			}
			lastResultDigit = lastResultDigit.next			
		}			
	}
	if carry != 0 {
		lastResultDigit.next = &SingleLinkedListNode{
			value : carry,
		}
	}

	return result
}


// Runtime : O(a) || O(b) . Could not be determined as we don't have relationship between A and B
// Memory : O(a) || O(b) . Could not be determined as we don't have relationship between A and B
func sumReversed(a,b *SingleLinkedListNode) *SingleLinkedListNode {
	stackA := stack.New()
	stackB := stack.New()
	firstValInserted := false
	result := &SingleLinkedListNode {
		value : 0,
	}
	carry := 0

	for temp := a ; temp != nil ; temp = temp.next {
		stackA.Push(temp.value)
	}

	for temp := b ; temp != nil ; temp = temp.next {
		stackB.Push(temp.value)
	}

	for ; stackA.Len() > 0 || stackB.Len() > 0 ; {
		newDigit := carry
		if stackA.Len() > 0 {
			val := stackA.Pop().(int)
			newDigit += val
		}
		if stackB.Len() > 0{
			val := stackB.Pop().(int)
			newDigit += val
		}
		carry = newDigit / 10
		newDigit = newDigit % 10
		if !firstValInserted {
			result.value = newDigit
			firstValInserted = true
		} else {
			result = &SingleLinkedListNode{
				value : newDigit,
				next : result,			
			}		
		}			
	}

	if carry != 0 {
		result = &SingleLinkedListNode{
			value : carry,
			next : result,	
		}	
	}

	return result
}

func main(){
	test := func(a,b *SingleLinkedListNode, summer func(a,b *SingleLinkedListNode) *SingleLinkedListNode){
		fmt.Println("--------------")
		fmt.Printf("Summing ")
		a.printList()
		fmt.Printf("With ")
		b.printList()
		res := summer(a,b)
		fmt.Printf("Result ")
		res.printList()
	}
	fmt.Println("\n\nTesting normal")
	test(createSingleLinkedList(7,1,6),createSingleLinkedList(5,9,8),sum)
	test(createSingleLinkedList(7,8),createSingleLinkedList(5,9,5),sum)
	fmt.Println("\n\nTesting reversed")
	test(createSingleLinkedList(6,1,7),createSingleLinkedList(8,9,5),sumReversed)
	test(createSingleLinkedList(8,7),createSingleLinkedList(5,9,5),sumReversed)
}
