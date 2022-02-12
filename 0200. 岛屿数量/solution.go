package leetcode

func dfs(grid [][]byte, i, j int, visited [][]bool) {
	m, n := len(grid), len(grid[0])
	//递归出口  越界 或者遇到海水
	if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == '0' {
		return
	}
	//递归出口  已经遍历过(i,j)
	if visited[i][j] {
		return
	}
	//遍历位置(i,j)
	visited[i][j] = true
	dfs(grid, i-1, j, visited)
	dfs(grid, i+1, j, visited)
	dfs(grid, i, j-1, visited)
	dfs(grid, i, j+1, visited)
}

func numIslands(grid [][]byte) int {
	res := 0
	m, n := len(grid), len(grid[0])
	//二维切片的声明方法
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	//遍历二维数组
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//满足条件
			if grid[i][j] == '1' && visited[i][j] == false {
				res++
				dfs(grid, i, j, visited)
			}
		}
	}
	return res
}

func numIslands2(grid [][]byte) int {
	res := 0
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		//递归出口  越界 或者遇到海水
		if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == '0' {
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
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//满足条件
			if grid[i][j] == '1' && visited[i][j] == false {
				res++
				dfs(i, j)
			}
		}
	}
	return res
}