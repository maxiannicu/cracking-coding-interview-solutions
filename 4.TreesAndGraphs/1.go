package main

import "fmt"

func visit(now, dest *GraphNode, visitedNodes map[*GraphNode]bool) bool {
	if now == dest {
		return true
	}
	if _, ok := visitedNodes[now]; ok {
		return false
	}

	result := false

	visitedNodes[now] = true

	if now.connections != nil {
		for _, conn := range now.connections {
			result = result || visit(conn, dest, visitedNodes)
			if result {
				return true
			}
		}
	}

	delete(visitedNodes, now)

	return result
}

func isRoute(from, to *GraphNode) bool {
	visitedNodes := map[*GraphNode]bool{}

	return visit(from, to, visitedNodes)
}

func main() {
	a := &GraphNode{
		name: "A",
	}
	b := &GraphNode{
		name: "B",
	}
	c := &GraphNode{
		name: "C",
	}
	d := &GraphNode{
		name: "D",
	}
	e := &GraphNode{
		name: "E",
	}
	f := &GraphNode{
		name: "F",
	}
	a.Connect(b, c, d)
	b.Connect(e)
	f.Connect(e, c, d)

	fmt.Printf("Is route between %v and %v ? %v\n", a, e, isRoute(a, e))
	fmt.Printf("Is route between %v and %v ? %v\n", a, f, isRoute(a, f))
}
