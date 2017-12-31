package main

import "fmt"

func removeDuplicates(head *SingleLinkedListNode){
	alreadyFoundElements := map[string]bool{}
	
	var last *SingleLinkedListNode
	for ; head != nil; head = head.next {
		if last != nil{
			if _,ok := alreadyFoundElements[head.value]; ok {
				last.next = head.next
			}
		}
		alreadyFoundElements[head.value] = true
		last = head
	}
}


func removeDuplicatesWithoutBuffer(head *SingleLinkedListNode){
	current := head
	
	var last *SingleLinkedListNode
	for ; current != nil; current = current.next {
		if last != nil{
			duplicated := false
			for temp := head; temp != nil && temp != current; temp = temp.next{
				if temp.value == current.value {
					duplicated = true
				}
			}
			if duplicated {
				last.next = head.next
			}
		}		
		last = current
	}
}


func main(){
	checkRemove := func(name string,duplicateRemover func(head *SingleLinkedListNode)){
		fmt.Println(name)
		fmt.Println("--------------------")
		list := getNewSingleLinkedList()
		fmt.Println("Before:")
		list.printList()
		removeDuplicates(list)
		fmt.Println("After:")
		list.printList()
		fmt.Println("--------------------")
		fmt.Println()
	}
	
	checkRemove("With extra memory",removeDuplicates)
	checkRemove("Without extra memory",removeDuplicatesWithoutBuffer)
}