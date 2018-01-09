package main

import "fmt"


// Runtime : O(n)
// Memory : O(1)
func getIntersection(node,intersect *SingleLinkedListNode) *SingleLinkedListNode {
	lengthDiff := node.Length() - intersect.Length()
	head := node

	for i := 0 ; i < lengthDiff ; i++ {
		head = head.next
	}

	if head == intersect {
		return head
	}

	return nil
}

func main(){
	a := createSingleLinkedList(1,2,3,4)
	b := createSingleLinkedList(2,3,4)
	fmt.Printf("List A : ")
	a.printList()
	fmt.Printf("List B : ")
	b.printList()
	fmt.Println("A intersect B\t",getIntersection(a,b))
	fmt.Println("A intersect A[1:]\t",getIntersection(a,a.next))
	fmt.Println("A intersect A\t",getIntersection(a,a))
}
