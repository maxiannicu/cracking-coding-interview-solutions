package main

import "fmt"

// Runtime : O(n)
// Memory : O(1)
func findElementFromTail(head *SingleLinkedListNode, kth int) *SingleLinkedListNode {
	pOne := head
	pTwo := head
	
	for i := 0;i < kth; i++ {
		if pOne == nil {
			return nil
		}
		pOne = pOne.next
	}
	
	for ; pOne != nil ; {
		pOne = pOne.next
		pTwo = pTwo.next
	}
	
	return pTwo
}


func main(){
	list := getNewSingleLinkedList()
	list.printList()
	if res := findElementFromTail(list, 5); res != nil {	
		fmt.Println(res.value)
	} else {
		fmt.Println("Not found")
	}
}