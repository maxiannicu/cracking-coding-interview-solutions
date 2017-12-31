package main

import "fmt"

type SingleLinkedListNode struct {
	value int
	next *SingleLinkedListNode
}

func getNewSingleLinkedList() *SingleLinkedListNode {
	return &SingleLinkedListNode{
		value : 1,
		next : &SingleLinkedListNode{
			value : 4,
			next : &SingleLinkedListNode{
				value : 8,
				next : &SingleLinkedListNode {
					value : 4,
					next : &SingleLinkedListNode {
						value : 2,
						next : &SingleLinkedListNode {
							value : 6,
							next : &SingleLinkedListNode {
								value : 5,
								next : nil,
							},
						},
					},
				},
			},
		},	
	}
}

func (head *SingleLinkedListNode) printList() {
	for ; head != nil; head = head.next {
		fmt.Printf("%d ",head.value)
	}
	fmt.Println()
}