package main

//0 ms

// time  : O(rc)
// space : O(1) -> not using visited

// group of island
func numIslands2(grid [][]byte) int {
	islandCount := 0

	for i := range grid {
		for j, gc := range grid[i] {
			if gc == 48 {
				continue
			}
			explore2(i, j, grid)
			islandCount++
		}
	}

	return islandCount
}

func explore2(i, j int, grid [][]byte) {
	// boundary check
	if (i < 0 || i >= len(grid)) || (j < 0 || j >= len(grid[0]) || grid[i][j] == 48) {
		return
	}

	// converting visited node to 0 // water
	grid[i][j] = 48

	// DFS
	// getting the neigh
	explore2(i-1, j, grid)
	explore2(i+1, j, grid)
	explore2(i, j-1, grid)
	explore2(i, j+1, grid)
}
