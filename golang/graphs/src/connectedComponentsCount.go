package main

// find total number of components (connect graphs)

// undirected
func connectedComponents(graph Graph) int {
	count := 0
	// act as a set
	visited := make(map[string]struct{})

	var stack Stack

	for k := range graph.Map {
		if _, ok := visited[k]; ok {
			continue
		}

		stack.Push(k)
		count++

		for !stack.IsEmpty() {
			poped := stack.Pop()
			visited[poped] = struct{}{}
			neigh := graph.Map[poped]
			for _, n := range neigh {
				if _, ok := visited[n]; !ok {
					stack.Push(n)
				}
			}
		}

	}

	return count
}
