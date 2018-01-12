package main

import "fmt"

type GraphNode struct {
	name        string
	connections []*GraphNode
}

func (this *GraphNode) Connect(nodes ...*GraphNode) {
	if this.connections == nil {
		this.connections = make([]*GraphNode, 0)
	}

	for _, conn := range nodes {
		this.connections = append(this.connections, conn)
	}
}

func (this *GraphNode) String() string {
	return fmt.Sprintf("%s", this.name)
}
