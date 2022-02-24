#### [1905. 统计子岛屿](https://leetcode-cn.com/problems/count-sub-islands/)

> 难度中等

给你两个 `m x n` 的二进制矩阵 `grid1` 和 `grid2` ，它们只包含 `0` （表示水域）和 `1` （表示陆地）。一个 **岛屿** 是由 **四个方向** （水平或者竖直）上相邻的 `1` 组成的区域。任何矩阵以外的区域都视为水域。

如果 `grid2` 的一个岛屿，被 `grid1` 的一个岛屿 **完全** 包含，也就是说 `grid2` 中该岛屿的每一个格子都被 `grid1` 中同一个岛屿完全包含，那么我们称 `grid2` 中的这个岛屿为 **子岛屿** 。

请你返回 `grid2` 中 **子岛屿** 的 **数目** 。

**示例 1：**

![img](https://assets.leetcode.com/uploads/2021/06/10/test1.png)

```
输入：grid1 = [[1,1,1,0,0],[0,1,1,1,1],[0,0,0,0,0],[1,0,0,0,0],[1,1,0,1,1]], grid2 = [[1,1,1,0,0],[0,0,1,1,1],[0,1,0,0,0],[1,0,1,1,0],[0,1,0,1,0]]
输出：3
解释：如上图所示，左边为 grid1 ，右边为 grid2 。
grid2 中标红的 1 区域是子岛屿，总共有 3 个子岛屿。
```

**示例 2：**

![img](https://assets.leetcode.com/uploads/2021/06/03/testcasex2.png)

```
输入：grid1 = [[1,0,1,0,1],[1,1,1,1,1],[0,0,0,0,0],[1,1,1,1,1],[1,0,1,0,1]], grid2 = [[0,0,0,0,0],[1,1,1,1,1],[0,1,0,1,0],[0,1,0,1,0],[1,0,0,0,1]]
输出：2 
解释：如上图所示，左边为 grid1 ，右边为 grid2 。
grid2 中标红的 1 区域是子岛屿，总共有 2 个子岛屿。 
```

**提示：**

- `m == grid1.length == grid2.length`
- `n == grid1[i].length == grid2[i].length`
- `1 <= m, n <= 500`
- `grid1[i][j]` 和 `grid2[i][j]` 都要么是 `0` 要么是 `1` 

#### 解题思路

这道题乍一看可以使用岛屿模板dfs来解决，但**难点**在于：如何快速判断 grid2 中是不是 grid1 的子岛屿。根据题意如果 `grid2` 中该岛屿的每一个格子都被 `grid1` 中同一个岛屿完全包含，那么我们称 `grid2` 中的这个岛屿为 **子岛屿** 。这句话也可以理解为如果在 grid2中某个格子是陆地1，但在grid1中是海水0，那么这个格子构成的岛屿一定不是子岛屿。

我们可以根据这点先将 grid2中构不成子岛屿的岛屿排除掉，那么剩下的不就是子岛屿了吗。这道题很像**封闭式岛屿**，只不过封闭岛屿排除的是四条边上的岛屿，而本题是排除不可能构成子岛屿的岛屿。

#### 代码

```go
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
```

