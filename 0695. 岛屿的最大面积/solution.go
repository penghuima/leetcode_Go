package leetcode

//1土地 0 水
func maxAreaOfIsland(grid [][]int) int {
	res := 0
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) (area int) {
		//递归出口  越界 或者遇到海水
		if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == 0 {
			return 0
		}
		//递归出口  已经遍历过(i,j)
		if visited[i][j] {
			return 0
		}
		//遍历位置(i,j)
		visited[i][j] = true
		//本层逻辑：拿到所有需要的底层数据后，进行逻辑处理，返回给上一层
		return dfs(i-1, j) + dfs(i+1, j) + dfs(i, j-1) + dfs(i, j+1) + 1
	}
	//遍历二维数组
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//满足条件
			if grid[i][j] == 1 && visited[i][j] == false {
				res = max(res, dfs(i, j))
			}
		}
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
