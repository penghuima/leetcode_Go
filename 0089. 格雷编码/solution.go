package leetcode

func grayCode(n int) []int {
	res := make([]int, 1<<n)
	var path []int
	used := make(map[int]bool, n)
	var backTrack func() bool
	//格雷码以0开始
	used[0] = true
	path = append(path, 0)
	backTrack = func() bool {
		//如果结果切片的长度达到 2的n次方
		if len(path) == 1<<n {
			copy(res, path)
			return true //找到正确解不再递归
		}
		for i := 0; i < n; i++ {
			//1<<i 等于2的i次方
			temp := 1<<i ^ path[len(path)-1]
			if used[temp] {
				continue
			}
			used[temp] = true
			path = append(path, temp)
			flag := backTrack()
			if flag {
				//一层层向上返回false
				return true
			} else {
				used[temp] = false
				path = path[:len(path)-1]
			}
		}
		return false
	}
	backTrack()
	return res
}
