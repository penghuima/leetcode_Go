package leetcode

func combine(n int, k int) [][]int {
	var res [][]int
	var path []int

	var backTrack func(int)
	backTrack = func(start int) {
		//符合条件
		if len(path) == k {
			temp := make([]int, k)
			copy(temp, path)
			res = append(res, temp)
		}
		for i := start; i <= n; i++ { //i从1开始
			//做出选择
			path = append(path, i)
			backTrack(i + 1)
			//回溯
			path = path[:len(path)-1]
		}
	}
	backTrack(1)
	return res
}