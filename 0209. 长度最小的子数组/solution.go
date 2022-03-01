package leetcode

import "math"

func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	sum := 0
	ans := math.MaxInt64
	for right < len(nums) {
		//先将元素添加进来窗口     right指针右边移动
		sum += nums[right]
		right++
		//满足缩减窗口的条件   和>=target    left指针左边移动
		for sum >= target {
			ans = min(ans, right-left)
			sum -= nums[left]
			left++
		}
	}
	if ans == math.MaxInt64 {
		return 0
	}
	return ans
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
