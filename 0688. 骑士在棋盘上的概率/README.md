#### [688. 骑士在棋盘上的概率](https://leetcode-cn.com/problems/knight-probability-in-chessboard/)

> 难度中等

在一个 `n x n` 的国际象棋棋盘上，一个骑士从单元格 `(row, column)` 开始，并尝试进行 `k` 次移动。行和列是 **从 0 开始** 的，所以左上单元格是 `(0,0)` ，右下单元格是 `(n - 1, n - 1)` 。

象棋骑士有8种可能的走法，如下图所示。每次移动在基本方向上是两个单元格，然后在正交方向上是一个单元格。

![img](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/12/knight.png)

每次骑士要移动时，它都会随机从8种可能的移动中选择一种(即使棋子会离开棋盘)，然后移动到那里。

骑士继续移动，直到它走了 `k` 步或离开了棋盘。

返回 *骑士在棋盘停止移动后仍留在棋盘上的概率* 。

**示例 1：**

```
输入: n = 3, k = 2, row = 0, column = 0
输出: 0.0625
解释: 有两步(到(1,2)，(2,1))可以让骑士留在棋盘上。
在每一个位置上，也有两种移动可以让骑士留在棋盘上。
骑士留在棋盘上的总概率是0.0625。
```

**示例 2：**

```
输入: n = 1, k = 0, row = 0, column = 0
输出: 1.00000
```

**提示:**

- `1 <= n <= 25`
- `0 <= k <= 100`
- `0 <= row, column <= n`

#### 解题思路

每一个骑士一共8种走法

- 令`dp[step][i][j]`表示从从棋盘上的点 (i, j)出发，走了 step 步时仍然留在棋盘上的概率
- 初始化，当（i，j）在棋盘上时 `dp[0][i][j]`=1 ；当（i，j）不在棋盘上时 `dp[step][i][j]`=0
- 递推 `dp[step][i][j]+=dp[step-1][x][y]/8（遍历8种走法）` 
- 模拟

#### 代码

```go
var dir = [][]int{{-2, -1}, {-2, 1}, {2, -1}, {2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}}

func knightProbability(n int, k int, row int, column int) float64 {
	//初始化
	dp := make([][][]float64, k+1)
	for step, _ := range dp {
		dp[step] = make([][]float64, n)
		for i := 0; i < n; i++ {
			dp[step][i] = make([]float64, n)
			for j := 0; j < n; j++ {
				//第一轮，只要在棋盘上的肯定都为1 啊
				if step == 0 {
					dp[step][i][j] = 1
				} else {
					for _, d := range dir {
						x, y := i+d[0], j+d[1]
						//还在棋盘上
						if x >= 0 && x < n && y >= 0 && y < n {
							dp[step][i][j] += dp[step-1][x][y] / 8
						}
					}
				}
			}

		}
	}
	return dp[k][row][column]
}
```

