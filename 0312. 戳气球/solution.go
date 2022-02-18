package leetcode

//戳气球 回溯暴力解决 [7,9,8,0,7,1,3,5,5,2,3]就超时了 执行耗时4s多
func maxCoins1(nums []int) int {
	res := 0
	sum := 0
	var backTrack func()
	backTrack = func() {
		//如果满足结束条件
		if len(nums) == 0 {
			res = max(res, sum) //记录最大值
			return
		}
		//回溯
		for i, v := range nums {
			//做选择
			temp := 0
			//计算乘积  nums[i-1] * v * nums[i+1] 考虑边界
			if i-1 < 0 {
				if i+1 > len(nums)-1 {
					temp = v
				} else {
					temp = v * nums[i+1]
				}
			} else if i+1 > len(nums)-1 {
				temp = nums[i-1] * v
			} else {
				temp = nums[i-1] * v * nums[i+1]
			}
			sum += temp
			//删除切片中的v
			nums = append(nums[:i], nums[i+1:]...)
			backTrack()
			//撤销选择
			sum -= temp
			//将v再加回到切片指定位置
			nums = append(nums[:i], append([]int{v}, nums[i:]...)...)
		}
	}
	backTrack()
	return res
}
//戳气球 动态规划
func maxCoins(nums []int) int {
	//将数组头尾添加元素1
	nums = append([]int{1}, nums[:]...)
	nums = append(nums, 1)
	dp := make([][]int, len(nums))
	for i, _ := range dp {
		dp[i] = make([]int, len(nums))
	}
	for i := len(nums) - 1; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+nums[i]*nums[k]*nums[j])
			}
		}
	}
	return dp[0][len(nums)-1]
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
//动态规划
