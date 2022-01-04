package leetcode

func permute(nums []int) [][]int {
	var res [][]int                  //最终要返回的二维数组切片
	var path []int                   //记录路径
	var visited = make(map[int]bool) //对已经做过的选择进行标记
	length := len(nums)
	var backTrack func()
	backTrack = func() {
		//如果满足结束条件，则将路径加入到二维数组切片中
		if len(path) == length {
			temp := make([]int, length)
			copy(temp, path)
			res = append(res, temp)
			return
		}
		//回溯
		for i := 0; i < length; i++ {
			//做选择
			if visited[nums[i]] == true {
				continue
			}
			path = append(path, nums[i])
			visited[nums[i]] = true
			backTrack()
			//撤销选择
			path = path[:len(path)-1]
			visited[nums[i]] = false
		}
	}
	backTrack()
	return res
}
