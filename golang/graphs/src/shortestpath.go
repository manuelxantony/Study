package main

import (
	"fmt"
	"graph/internal"
)

// {"i", "j"},
// {"k", "i"},
// {"m", "k"},
// {"k", "l"},
// {"o", "n"},

func shortedPath(edges [][]string, src, dst string) int {
	graph := buildGraphHere(edges)
	visited := make(map[string]struct{})
	// BFS
	queue := VDQueue{}
	srcNode := internal.QueueNode{
		Value: src,
	}
	queue.Enqueue(internal.QueueNode(srcNode))
	visited[src] = struct{}{}

	for !queue.IsEmpty() {
		poped := queue.Dequeue()
		if poped.Value == dst {
			return poped.Distance
		}
		for _, neigh := range graph.Map[poped.Value] {
			if _, ok := visited[neigh]; !ok {
				neighNode := internal.QueueNode{
					Value:    neigh,
					Distance: poped.Distance + 1,
				}
				queue.Enqueue(neighNode)
				visited[neigh] = struct{}{}
			}
		}
	}

	return -1
}

func buildGraphHere(edges [][]string) Graph {
	graph := Graph{
		Map: make(map[string][]string),
	}

	for _, node := range edges {
		graph.Map[node[0]] = append(graph.Map[node[0]], node[1])
		graph.Map[node[1]] = append(graph.Map[node[1]], node[0])
	}
	fmt.Println(graph)
	return graph
}
