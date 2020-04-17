package main

func bfs(grid [][]byte, i, j int) {

	// grid boundary check
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'

	// up, down, left, right
	bfs(grid, i+1, j)
	bfs(grid, i-1, j)
	bfs(grid, i, j+1)
	bfs(grid, i, j-1)

}

func numIslands(grid [][]byte) int {
	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' { // new island found
				count++
				bfs(grid, i, j)
			}
		}
	}
	return count

}
