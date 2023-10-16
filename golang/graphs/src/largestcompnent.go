package main

import "fmt"

// undirected
func largestComponent(graph Graph) int {
	var count int
	var maxCount int
	// act as a set
	//visited := make(map[string]struct{})
	var stack Stack

	for k := range graph.Map {
		visited := make(map[string]struct{})
		stack.Push(k)
		visited[k] = struct{}{}
		for !stack.IsEmpty() {
			poped := stack.Pop()
			neigh := graph.Map[poped]
			for _, n := range neigh {
				if _, ok := visited[n]; !ok {
					stack.Push(n)
					visited[n] = struct{}{}
				}
			}
		}
		count = len(visited)

		if count > maxCount {
			fmt.Println(count)
			maxCount = count
		}

	}

	return maxCount
}
