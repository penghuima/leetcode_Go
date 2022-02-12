#### [1020. 飞地的数量](https://leetcode-cn.com/problems/number-of-enclaves/)

> 难度中等

给你一个大小为 `m x n` 的二进制矩阵 `grid` ，其中 `0` 表示一个海洋单元格、`1` 表示一个陆地单元格。

一次 **移动** 是指从一个陆地单元格走到另一个相邻（**上、下、左、右**）的陆地单元格或跨过 `grid` 的边界。

返回网格中 **无法** 在任意次数的移动中离开网格边界的陆地单元格的数量。

**示例 1：**

![img](https://assets.leetcode.com/uploads/2021/02/18/enclaves1.jpg)

```
输入：grid = [[0,0,0,0],[1,0,1,0],[0,1,1,0],[0,0,0,0]]
输出：3
解释：有三个 1 被 0 包围。一个 1 没有被包围，因为它在边界上。
```

**示例 2：**

![img](https://assets.leetcode.com/uploads/2021/02/18/enclaves2.jpg)

```
输入：grid = [[0,1,1,0],[0,0,1,0],[0,0,1,0],[0,0,0,0]]
输出：0
解释：所有 1 都在边界上或可以到达边界。
```

**提示：**

- `m == grid.length`
- `n == grid[i].length`
- `1 <= m, n <= 500`
- `grid[i][j]` 的值为 `0` 或 `1`

#### 解题思路

岛屿问题

这道题就是[统计岛屿](https://leetcode-cn.com/problems/number-of-islands/)的变形，我们要[统计封闭岛屿](https://leetcode-cn.com/problems/number-of-closed-islands/)的方格大小。我们首先将矩阵的四个边用dfs算法给跑一遍，那么再遍历矩阵，只要其

`grid[i][j]==1 &&  visited[i][j] == false` 就一定是封闭岛屿所属的方格

#### 代码

```go
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
```

