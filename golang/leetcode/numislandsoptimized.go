package main

// time  : O(rc)
// space : O(1) -> not using visited

// group of island
func numIslands3(grid [][]byte) int {
	islandCount := 0

	for i := range grid {
		for j, gc := range grid[i] {
			if gc == 48 {
				continue
			}
			explore3(i, j, grid)
			islandCount++
		}
	}

	return islandCount
}

func explore3(i, j int, grid [][]byte) {
	// boundary check
	if (i < 0 || i >= len(grid)) || (j < 0 || j >= len(grid[0]) || grid[i][j] == 48) {
		return
	}

	// converting visited node to 0 // water
	grid[i][j] = 48

	// DFS
	// getting the neigh
	explore3(i-1, j, grid)
	explore3(i+1, j, grid)
	explore3(i, j-1, grid)
	explore3(i, j+1, grid)
}
