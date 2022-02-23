#### [1254. 统计封闭岛屿的数目](https://leetcode-cn.com/problems/number-of-closed-islands/)

> 难度中等

二维矩阵 `grid` 由 `0` （土地）和 `1` （水）组成。岛是由最大的4个方向连通的 `0` 组成的群，封闭岛是一个 `完全` 由1包围（左、上、右、下）的岛。

请返回 *封闭岛屿* 的数目。

**示例 1：**

![img](https://assets.leetcode.com/uploads/2019/10/31/sample_3_1610.png)

```
输入：grid = [[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,0,1,0,1],[1,1,1,1,1,1,1,0]]
输出：2
解释：
灰色区域的岛屿是封闭岛屿，因为这座岛屿完全被水域包围（即被 1 区域包围）。
```

**示例 2：**

![img](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/11/07/sample_4_1610.png)

```
输入：grid = [[0,0,1,0,0],[0,1,0,1,0],[0,1,1,1,0]]
输出：1
```

**示例 3：**

```
输入：grid = [[1,1,1,1,1,1,1],
             [1,0,0,0,0,0,1],
             [1,0,1,1,1,0,1],
             [1,0,1,0,1,0,1],
             [1,0,1,1,1,0,1],
             [1,0,0,0,0,0,1],
             [1,1,1,1,1,1,1]]
输出：2
```

**提示：**

- `1 <= grid.length, grid[0].length <= 100`
- `0 <= grid[i][j] <=1`

#### 解题思路

这个题和统计岛屿的数量差不多，唯一的区别就是这里统计的岛屿是封闭的，处于矩阵四周边上的岛屿不算。因此统计岛屿的时候我们可以先将矩阵四周用dfs算法跑一遍，使visited对应位置为 true ,然后就是再遍历整个矩阵即可

> 错误做法：在遍历矩阵的时候直接从[1,m-1] [1,n-1] 遍历，这样岛屿会计算错误多的答案

#### 代码

```go
//0 表示陆地，用 1 表示海水：
func closedIsland(grid [][]int) int {
	res := 0
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		//递归出口  越界 或者遇到海水
		if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == 1 {
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
		dfs(0, j)      //最上面
		dfs(m-1, j)    //最下面
	}
	for i := 0; i < m; i++ {
		dfs(i, 0)      //最左边
		dfs(i, n-1)    //最右边
	}
    //遍历二维数组
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//满足条件
			if grid[i][j] == 0 && visited[i][j] == false {
				res++
				dfs(i, j)
			}
		}
	}
	return res
}
```

