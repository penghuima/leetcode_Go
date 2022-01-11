package leetcode

func subsets(nums []int) [][]int {
	var res [][]int
	var path []int
	length := len(nums)
	var backTrack func(int)
	backTrack = func(start int) {
		//满足结束条件
		temp := make([]int, len(path)) //为啥使用make就正确，如果使用变量 var temp1 []int,就输出错误呢
		copy(temp, path)
		res = append(res, temp)
		for i := start; i < length; i++ {
			//做选择
			path = append(path, nums[i])
			backTrack(i + 1) //不识别在函数里传入 i++参数
			//撤销选择
			path = path[:len(path)-1]
		}
	}
	backTrack(0)
	return res
}