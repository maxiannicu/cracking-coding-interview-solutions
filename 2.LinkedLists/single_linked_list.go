package main

import "fmt"

type SingleLinkedListNode struct {
	value int
	next *SingleLinkedListNode
}

func (node *SingleLinkedListNode) String() string {
	return fmt.Sprintf("Node(value : %v, next : %v)", node.value, node.next != nil)
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

func createSingleLinkedList(elements ...int) *SingleLinkedListNode {
	var head *SingleLinkedListNode
	var last *SingleLinkedListNode

	for _, el := range elements {
		if head == nil {
			head = &SingleLinkedListNode {
				value : el,
			}
			last = head
		} else {
			last.next = &SingleLinkedListNode {
				value : el,
			}
			last = last.next
		}
	}

	return head
}

func (head *SingleLinkedListNode) printList() {
	for ; head != nil; head = head.next {
		fmt.Printf("%d ",head.value)
	}
	fmt.Println()
}