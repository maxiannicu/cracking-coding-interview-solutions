package main

import "fmt"

func getLast(head *SingleLinkedListNode) *SingleLinkedListNode {
	temp := head 
	for ; temp.next != nil; temp = temp.next {}

	return temp
}


// Runtime : O(n)
// Memory : O(1)
func partition(val int, head *SingleLinkedListNode) *SingleLinkedListNode {
	var left,right *SingleLinkedListNode

	// Find the left and right sequences. One of them is used only for append, another for extraction
	if head.value >= val {
		right = head
		temp := head
		for ; temp.next != nil && temp.next.value >= val ; temp = temp.next {}
		left = temp.next
		temp.next = nil
	} else {
		left = head
		temp := head
		for ; temp.next != nil && temp.next.value < val ; temp = temp.next {}
		right = temp.next
		temp.next = nil
	}

	if left == nil { // if all items are greater or equal to partition
		return right
	} else if right == nil {  // if all items are less or equal to partition
 		return left
	} else {
		if head == left {
			lastLeft := getLast(left)
			var lastRight *SingleLinkedListNode

			// iterate over right sequence and move smaller items to left sequence
			for temp := right; temp != nil; temp = temp.next {
				if temp.value < val {
					lastRight.next = temp.next
					lastLeft.next = temp
					temp.next = nil
					lastLeft = lastLeft.next
					temp = lastRight
				} else {
					lastRight = temp
				}
			}
			lastLeft.next = right
		} else {
			lastRight := getLast(right)
			var lastLeft *SingleLinkedListNode

			// iterate over left seuqence and move greater or equal items to right sequence
			for temp := left; temp != nil; temp = temp.next {
				if temp.value >= val {
					lastLeft.next = temp.next
					lastRight.next = temp
					temp.next = nil
					lastRight = lastRight.next
					temp = lastLeft
				} else {
					lastLeft = temp
				}
			}
			lastLeft.next = right
		}
	}

	return left
}

func main(){
	testInputs := []*SingleLinkedListNode{
		createSingleLinkedList(1,5,3,2,6),
		createSingleLinkedList(5,1,3,2,6),
		createSingleLinkedList(8,7,6,5,4,3,2,1,0),
		createSingleLinkedList(1,2,0,-3,2,0),
	}
	partitionOver := 4
	for _, el := range testInputs {	
		fmt.Println("--------------")
		fmt.Println("The list is")
		el.printList()
		fmt.Println("After partition list is")
		head := partition(partitionOver,el)
		head.printList()
	}
}
