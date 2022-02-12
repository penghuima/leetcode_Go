#### [200. 岛屿数量](https://leetcode-cn.com/problems/number-of-islands/)

> 难度中等

给你一个由 `'1'`（陆地）和 `'0'`（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。

**示例 1：**

```
输入：grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
输出：1
```

**示例 2：**

```
输入：grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
输出：3
```

**提示：**

- `m == grid.length`
- `n == grid[i].length`
- `1 <= m, n <= 300`
- `grid[i][j]` 的值为 `'0'` 或 `'1'`

#### 解题思路

知道使用dfs解答，并也知道dfs的框架

难点：如何标记岛屿

```go
//遍历数组，当访问到1且没被访问过的时候，岛屿数量加1
for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//满足条件
			if grid[i][j] == '1' && visited[i][j] == false {
				res++
				dfs(i, j)
			}
		}
	}
```

#### 代码

//代码1   定义函数变量

```go
func numIslands(grid [][]byte) int {
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
```

//代码2   

```go
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
```

