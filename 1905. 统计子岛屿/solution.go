package leetcode

//1 陆地 0海洋
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	res := 0
	for i, g1 := range grid1 {
		for j, v1 := range g1 {
			//一定不可能构成子岛屿
			if v1 == 0 && grid2[i][j] == 1 {
				dfs(grid2, i, j)
			}
		}
	}
	//统计剩余的岛屿
	for i, g2 := range grid2 {
		for j, v2 := range g2 {
			if v2 == 1 {
				res++
				dfs(grid2, i, j)
			}
		}
	}
	return res
}
func dfs(grid [][]int, i, j int) {
	m, n := len(grid), len(grid[0])
	if i < 0 || j < 0 || i >= m || j >= n {
		return
	}
	if grid[i][j] == 0 {
		return
	}
	grid[i][j] = 0
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}
