#### [463. 岛屿的周长](https://leetcode-cn.com/problems/island-perimeter/)

> 难度简单

给定一个 `row x col` 的二维网格地图 `grid` ，其中：`grid[i][j] = 1` 表示陆地， `grid[i][j] = 0` 表示水域。

网格中的格子 **水平和垂直** 方向相连（对角线方向不相连）。整个网格被水完全包围，但其中**恰好**有一个岛屿（或者说，一个或多个表示陆地的格子相连组成的岛屿）。

岛屿中没有“湖”（“湖” 指水域在岛屿内部且不和岛屿周围的水相连）。格子是边长为 1 的正方形。网格为长方形，且宽度和高度均不超过 100 。计算这个岛屿的周长。

**示例 1：**

![img](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/12/island.png)

```
输入：grid = [[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]
输出：16
解释：它的周长是上面图片中的 16 个黄色的边
```

**示例 2：**

```
输入：grid = [[1]]
输出：4
```

**示例 3：**

```
输入：grid = [[1,0]]
输出：4
```

**提示：**

- `row == grid.length`
- `col == grid[i].length`
- `1 <= row, col <= 100`
- `grid[i][j]` 为 `0` 或 `1`

#### 解题思路

这道题是计算岛屿的周长，显然如果没有格子交互的话，那么每个格子都应该为周长贡献 4 , 考虑每个格子上下左右是否有邻接的岛屿格子，每有一个邻近的格子，则其周长贡献-1，很显然四周被包围的格子贡献为0。

**扩展思考**

这个题设置的太简单了，每张二维图只有一个岛屿，如果将题目改为岛屿数目不定，让你返回每个岛屿的周长呢？

这时就需要将思路改写为dfs模板的样子，解题思路也改变了：对于一个陆地格子的每条边，它被算作岛屿的周长当且仅当**这条边为网格的边界或者相邻的另一个格子为水域**。 因此，我们可以遍历每个陆地格子，看其四个方向是否为边界或者水域，如果是，将这条边的贡献（即 1）

#### 代码

> 原题目

```go
var dir = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
func islandPerimeter(grid [][]int) int {
	res := 0
	for i, g := range grid {
		for j, v := range g {
			//
			if v == 1 {
				res += 4
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
	}
	return res
}
```

> 修改后的题目:将题目改为岛屿数目不定，让你返回每个岛屿的周长呢？

```go
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
```

