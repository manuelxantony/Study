package main

import "fmt"

func numIslands(grid [][]byte) int {
	count := 0
	// map will be for identifing colum and row
	// eg: for row ,col 3, 4 it will be as visited[34] = struct{}{}
	// eg: for row ,col 4, 3 it will be as visited[43] = struct{}{}
	visited := make(map[string]struct{})

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// or len of visited can be checked for inc count
			if explore(i, j, grid, visited) {
				count++
			}
		}
	}
	return count
}

func explore(i, j int, grid [][]byte, visited map[string]struct{}) bool {
	if (i < 0 || i > len(grid)-1) || (j < 0 || j > len(grid[0])-1) {
		return false
	}

	vistedGrid := fmt.Sprintf("%d%d", i, j)
	if _, ok := visited[vistedGrid]; ok {
		return false
	}

	if grid[i][j] == 0 {
		return false
	}

	if grid[i][j] == 1 {
		// DFS
		vistedGrid := fmt.Sprintf("%d%d", i, j)
		visited[vistedGrid] = struct{}{}

		// getting the neigh L
		explore(i-1, j, grid, visited)
		explore(i+1, j, grid, visited)
		explore(i, j-1, grid, visited)
		explore(i, j+1, grid, visited)
	}
	return true
}
