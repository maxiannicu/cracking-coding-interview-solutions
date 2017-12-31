package main

import "fmt"

type SingleLinkedListNode struct {
	value string
	next *SingleLinkedListNode
}

func getNewSingleLinkedList() *SingleLinkedListNode {
	return &SingleLinkedListNode{
		value : "A",
		next : &SingleLinkedListNode{
			value : "B",
			next : &SingleLinkedListNode{
				value : "B",
				next : &SingleLinkedListNode {
					value : "C",
					next : &SingleLinkedListNode {
						value : "A",
						next : nil,
					},
				},
			},
		},	
	}
}

func (head *SingleLinkedListNode) printList() {
	for ; head != nil; head = head.next {
		fmt.Printf("%s ",head.value)
	}
	fmt.Println()
}