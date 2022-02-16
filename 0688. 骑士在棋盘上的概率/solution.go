package leetcode

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