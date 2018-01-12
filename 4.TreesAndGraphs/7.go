package main

import "fmt"

func appendIfNotPresent(arr []*GraphNode, elements ...*GraphNode) []*GraphNode {
	for _, el := range elements {
		present := false
		for _, existingEl := range arr {
			if existingEl == el {
				present = true
			}
		}

		if !present {
			arr = append(arr, el)
		}
	}
	return arr
}

func build(node *GraphNode, currentBuilds map[*GraphNode]bool) ([]*GraphNode, bool) {
	if _, ok := currentBuilds[node]; ok {
		return nil, false
	}

	currentBuilds[node] = true

	aggregation := make([]*GraphNode, 0)

	for _, el := range node.connections {
		if res, ok := build(el, currentBuilds); ok {
			aggregation = append(aggregation, res...)
		} else {
			return nil, false
		}
	}

	delete(currentBuilds, node)

	return append(aggregation, node), true
}

func findOrder(nodes []*GraphNode) ([]*GraphNode, bool) {
	buildOrder := make([]*GraphNode, 0)

	for _, el := range nodes {
		currentBuilds := map[*GraphNode]bool{}
		if result, ok := build(el, currentBuilds); ok {
			buildOrder = appendIfNotPresent(buildOrder, result...)
		} else {
			return nil, false
		}
	}

	return buildOrder, true
}

func main() {
	a := &GraphNode{
		name: "a",
	}
	b := &GraphNode{
		name: "b",
	}
	c := &GraphNode{
		name: "c",
	}
	d := &GraphNode{
		name: "d",
	}
	e := &GraphNode{
		name: "e",
	}
	f := &GraphNode{
		name: "f",
	}
	d.Connect(a)
	b.Connect(f)
	d.Connect(b)
	a.Connect(f)
	c.Connect(d)
	projects := []*GraphNode{
		a, b, c, d, e, f,
	}

	if order, buildResult := findOrder(projects); buildResult {
		for _, el := range order {
			fmt.Println(el)
		}
	} else {
		fmt.Println("Could not build. Circular dependecy")
	}
}
