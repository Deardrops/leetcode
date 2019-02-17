package main

func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rowLen := len(grid)
	colLen := len(grid[0])
	count := 0
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			count += countIsland(i, j, grid)
		}
	}
	return count
}

func countIsland(i int, j int, grid [][]byte) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || string(grid[i][j]) != "1" {
		return 0
	}
	//fmt.Println(string(grid[i][j]))
	grid[i][j] = byte('0')
	countIsland(i+1, j, grid)
	countIsland(i-1, j, grid)
	countIsland(i, j+1, grid)
	countIsland(i, j-1, grid)
	return 1
}
