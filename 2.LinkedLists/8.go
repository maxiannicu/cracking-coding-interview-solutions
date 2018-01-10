package main

import "fmt"


// Runtime : O(n)
// Memory : O(1)
func getCycleNode(node *SingleLinkedListNode) *SingleLinkedListNode {
	slow,fast := node,node
	
	for {
		fast = fast.next.next
		slow = slow.next

		if slow == fast {
			break
		}
	}

	slow = node
	
	for ; slow != fast; {
		slow = slow.next
		fast = fast.next
	}

	return slow
}

func getLast(node *SingleLinkedListNode) *SingleLinkedListNode {
	temp := node
	for ; temp.next != nil; temp = temp.next {}

	return temp
}

func main(){
	a := createSingleLinkedList(1,2,3,4,5,6)
	last := getLast(a)
	last.next = a.next

	fmt.Println(getCycleNode(a))	
}
