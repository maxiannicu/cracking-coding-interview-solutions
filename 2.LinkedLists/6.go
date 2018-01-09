package main

import "fmt"
import "github.com/golang-collections/collections/stack"


// Runtime : O(n)
// Memory : O(n)
func isPalindrome(node *SingleLinkedListNode) bool {
	length := node.Length()
	s := stack.New()
	
	head := node
	for i:=0; i < length; i++ {
		if !(length % 2 == 1 && i == length/2) {
			if i < length/2 {
				s.Push(head.value)
			} else {
				val := s.Pop().(int)
				if val != head.value {
					return false
				}
			}
		}
		head = head.next
	}

	return true
}

func main(){
	testInputs := []*SingleLinkedListNode {
		createSingleLinkedList(1,2,1),
		createSingleLinkedList(1,2,2,1),
		createSingleLinkedList(1,2,3),
		createSingleLinkedList(1,2,2,3),
	}

	for _, el := range testInputs {
		el.printList()
		fmt.Printf(" is palindrome ? %v\n", isPalindrome(el))
	}
}
