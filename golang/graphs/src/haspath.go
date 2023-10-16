package main

func hasPath(g *Graph, src, dst string) bool {
	// dfs
	// var stack Stack
	// stack.Push(src)

	// for !stack.IsEmpty() {
	// 	neig := g.Map[stack.Pop()]
	// 	for _, n := range neig {
	// 		if n == dst {
	// 			return true
	// 		}
	// 		stack.Push(neig...)
	// 	}
	// }
	// return false

	//-------------------------------------------
	// dfs recursion
	// if src == dst {
	// 	return true
	// }

	// neig := g.Map[src]
	// for _, n := range neig {
	// 	if hasPath(g, n, dst) {
	// 		return true
	// 	}
	// }
	// return false

	//-------------------------------------------
	// bfs
	if src == dst {
		return true
	}
	var queue Queue
	queue.Enqueue(src)

	for !queue.IsEmpty() {
		neig := g.Map[queue.Dequeue()]
		for _, n := range neig {
			if n == dst {
				return true
			}
			queue.Enqueue(n)
		}
	}
	return false
}
