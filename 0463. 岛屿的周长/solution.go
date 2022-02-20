package leetcode

var dir = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func islandPerimeter(grid [][]int) int {
	res := 0
	for i, g := range grid {
		for j, v := range g {
			//
			if v == 1 {
				res += 4
				grid[i][j] = 0 //防止后面被重复多减
			}
			for k := 0; k < 4; k++ {
				newI, newJ := i+dir[k][0], j+dir[k][1]
				if newI < 0 || newJ < 0 || newI >= len(grid) || newJ >= len(grid[0]) {
					continue
				} else if grid[newI][newJ] == 1 {
					res--
				}
			}
		}
	}
	return res
}

func islandPerimeter1(grid [][]int) []int {
	var res []int
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		//
		if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == 0 {
			return 1
		}
		if visited[i][j] {
			return 0
		}
		visited[i][j] = true
		return dfs(i-1, j) + dfs(i+1, j) + dfs(i, j-1) + dfs(i, j+1)
	}
	//遍历二维数组
	for i, r := range grid {
		for j, v := range r {
			if v == 1 && !visited[i][j] {
				res = append(res, dfs(i, j))
			}
		}
	}
	return res
}
