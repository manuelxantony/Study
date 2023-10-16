package main

// time  : O(rc)
// space : O(1) -> not using visited

// group of island
func numIslands2(grid [][]string) int {
	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// or len of visited can be checked for inc count
			if explore2(i, j, grid) {
				count++
			}
		}
	}
	return count
}

func explore2(i, j int, grid [][]string) bool {
	if (i < 0 || i > len(grid)-1) || (j < 0 || j > len(grid[0])-1) {
		return false
	}

	if grid[i][j] == "0" {
		return false
	}

	if grid[i][j] == "1" {
		grid[i][j] = "0"
		// DFS
		// getting the neigh 1
		explore2(i-1, j, grid)
		explore2(i+1, j, grid)
		explore2(i, j-1, grid)
		explore2(i, j+1, grid)
	}
	return true
}
