package main


// Runtime : O(n)
// Memory : O(1)
func deleteMiddle(head *SingleLinkedListNode, node *SingleLinkedListNode) {
	last := head
		
	for ;last != nil && last.next != node; last = last.next {}
	
	if last != nil {
		last.next = node.next
	}
}


func main(){
	list := getNewSingleLinkedList()
	list.printList()	
	el := list.next.next
	deleteMiddle(list, el)
	list.printList()
}