package main

import "fmt"

func undirectedPathDitection(edges [][]string, src, dst string) bool {

	if src == dst {
		return true
	}

	graph := buildGraph(edges)
	fmt.Println(graph)

	// BFS
	var queue Queue
	queue.Enqueue(src)

	visited := make(map[string]struct{})

	for !queue.IsEmpty() {
		poped := queue.Dequeue()
		visited[poped] = struct{}{}

		neigh := graph.Map[poped]
		for _, n := range neigh {
			if dst == n {
				return true
			}
			if _, ok := visited[n]; !ok {
				queue.Enqueue(n)
			}
		}

	}

	return false
}

func buildGraph(edges [][]string) *Graph {
	g := &Graph{
		Map: make(map[string][]string),
	}
	for _, vertices := range edges {
		g.Map[vertices[0]] = append(g.Map[vertices[0]], vertices[1])
		g.Map[vertices[1]] = append(g.Map[vertices[1]], vertices[0])
	}
	return g
}
