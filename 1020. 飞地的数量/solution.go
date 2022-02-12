package leetcode

//0是海洋 1是陆地
func numEnclaves(grid [][]int) int {
	res := 0
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		//递归出口  越界 或者遇到海水
		if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == 0 {
			return
		}
		//递归出口  已经遍历过(i,j)
		if visited[i][j] {
			return
		}
		//遍历位置(i,j)
		visited[i][j] = true
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}
	//遍历二维数组
	//首先把四周给遍历了
	for j := 0; j < n; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}
	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//满足条件
			if grid[i][j] == 1 && visited[i][j] == false {
				res++
			}
		}
	}
	return res
}
